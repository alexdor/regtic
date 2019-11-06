package db

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type validationLocks struct {
	companies sync.Mutex
	people    sync.Mutex
	errors    sync.Mutex
}
type People struct {
	Bad     models.BadPersonSlice `json:"bad"`
	Warning models.BadPersonSlice `json:"warning"`
	Good    models.PersonSlice    `json:"good"`
}
type ValidationResponse struct {
	Info      models.Company      `json:"info"`
	Companies models.CompanySlice `json:"companies"`
	People    People              `json:"people"`
	Errors    []error             `json:"errors"`
}

func ValidateCompany(ctx context.Context, id string) (*ValidationResponse, error) {
	company, err := models.Companies(
		models.CompanyWhere.ID.EQ(id),
		qm.Load(models.CompanyRels.MotherCompanyCompanies),
		qm.Load(models.CompanyRels.Persons),
		qm.Limit(1),
	).One(ctx, DB)

	var response ValidationResponse
	if err != nil {
		return &response, err
	}

	locks := &validationLocks{}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	traverseThroughTheCompany(ctx, &models.CompanySlice{company}, &response, locks, wg)

	response.Info = *company

	wg.Wait()

	unique(response)
	return &response, nil
}

func unique(response ValidationResponse) {
	//TODO
	return
}

func traverseThroughTheCompany(ctx context.Context, companies *models.CompanySlice, response *ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	wg.Add(2)
	go getDaughterCompanies(ctx, companies, response, locks, wg)
	go getOwners(ctx, companies, response, locks, wg)
	wg.Done()
}

func getDaughterCompanies(ctx context.Context, companies *models.CompanySlice, response *ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()

	if companies == nil || len(*companies) == 0 {
		return
	}
	motherCompanies := models.CompanySlice{}
	for _, company := range *companies {
		mCompanies, err := company.MotherCompanyCompanies(qm.Load(models.CompanyRels.Persons), qm.Load(models.CompanyRels.MotherCompanyCompanies)).All(ctx, DB)
		motherCompanies = append(motherCompanies, mCompanies...)
		writeError(err, response, locks)
	}
	wg.Add(1)
	traverseThroughTheCompany(ctx, &motherCompanies, response, locks, wg)

	locks.companies.Lock()
	response.Companies = append(response.Companies, motherCompanies...)
	locks.companies.Unlock()
}

func getOwners(ctx context.Context, companies *models.CompanySlice, response *ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()
	if companies == nil || len(*companies) == 0 {
		return
	}

	for _, company := range *companies {
		persons, err := company.Persons().All(ctx, DB)
		if err != nil {
			writeError(err, response, locks)
			continue
		}
		wg.Add(1)
		go searchForBadPersons(ctx, &persons, response, locks, wg)
	}
}

var badPersonWhereClause = models.BadPersonColumns.NameVector + "@@ plainto_tsquery('simple', ? && ?)"

func searchForBadPersons(ctx context.Context, persons *models.PersonSlice, response *ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()
	if persons == nil || len(*persons) == 0 {
		return
	}
	peopleResponse := People{}
	for _, person := range *persons {
		//TODO: add aliases, score and sorting
		badPersons, err := models.BadPersons(
			qm.Where(badPersonWhereClause, person.FirstName, person.LastName),
		).All(ctx, DB)

		noRows := errors.As(err, &sql.ErrNoRows)
		if !noRows && err != nil {
			writeError(err, response, locks)
		}

		if noRows || len(badPersons) == 0 {
			peopleResponse.Good = append(peopleResponse.Good, person)
			continue
		}

		//TODO: Get all the matches and not just the first one
		switch badPersons[0].Type {
		case models.BadPersonTypePEP:
			peopleResponse.Warning = append(peopleResponse.Warning, badPersons[0])
		case models.BadPersonTypeSANCTION:
			peopleResponse.Bad = append(peopleResponse.Bad, badPersons[0])
		}
	}

	locks.people.Lock()
	response.People.Good = append(response.People.Good, peopleResponse.Good...)
	response.People.Warning = append(response.People.Warning, peopleResponse.Warning...)
	response.People.Bad = append(response.People.Bad, peopleResponse.Bad...)
	locks.people.Unlock()

}

func writeError(err error, response *ValidationResponse, locks *validationLocks) {
	if err != nil {
		locks.errors.Lock()
		response.Errors = append(response.Errors, err)
		locks.errors.Unlock()
	}
}
