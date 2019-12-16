package db

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/alexdor/regtic/api/interfaces"
	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type validationLocks struct {
	companiesMutex    sync.Mutex
	companiesMapMutex sync.RWMutex
	peopleMapMutex    sync.RWMutex
	peopleMutex       sync.Mutex
	errorsMutex       sync.Mutex
	companyMap        map[string]struct{}
	personMap         map[string]struct{}
}

func ValidateCompany(ctx context.Context, id string) (*interfaces.ValidationResponse, error) {
	company, err := models.Companies(
		models.CompanyWhere.ID.EQ(id),
		qm.Load(models.CompanyRels.MotherCompanyCompanyToCompanies, companyToCompanyFilter),
		qm.Load(models.CompanyRels.CompanyToPeople, companyToPeopleFilter),
		qm.Load(models.CompanyRels.BadCompanies),
	).One(ctx, DB)

	response := &interfaces.ValidationResponse{}
	if err != nil {
		return response, err
	}

	locks := &validationLocks{
		companyMap: make(map[string]struct{}),
		personMap:  make(map[string]struct{}),
	}
	locks.companyMap[company.ID] = struct{}{}

	wg := &sync.WaitGroup{}
	traverseThroughCompanies(ctx, &models.CompanySlice{company}, response, locks, wg)
	response.Info = DBCompanyToCompanyJson(company)
	jsonCompany, err := convertCompany(ctx, company)
	if err != nil {
		writeError(err, response, locks)
	} else {
		locks.companiesMutex.Lock()
		response.Companies = append(response.Companies, jsonCompany)
		locks.companiesMutex.Unlock()
	}
	wg.Wait()

	return response, nil
}

// Go through the company structure
func traverseThroughCompanies(ctx context.Context, companies *models.CompanySlice, response *interfaces.ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	wg.Add(2)
	go getMotherCompanies(ctx, companies, response, locks, wg)
	go getOwners(ctx, companies, response, locks, wg)
}

// Add mother companies to the list of companies
func getMotherCompanies(ctx context.Context, companies *models.CompanySlice, response *interfaces.ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()

	if companies == nil || len(*companies) == 0 {
		return
	}
	aggregatedMotherCompanies := models.CompanySlice{}

	for _, company := range *companies {

		motherCompanies, err := company.DaughterCompanyCompanyToCompanies(
			qm.Load(models.CompanyToCompanyRels.MotherCompany),
			companyToCompanyFilter,
		).All(ctx, DB)
		if err != nil {
			writeError(err, response, locks)
			continue
		}
		//Search for companies that have already been traversed and drop them from the list
		for i := range motherCompanies {
			locks.companiesMapMutex.RLock()
			_, parsedAlready := locks.companyMap[motherCompanies[i].MotherCompanyID]
			locks.companiesMapMutex.RUnlock()
			if parsedAlready {
				continue
			}
			locks.companiesMapMutex.Lock()
			locks.companyMap[motherCompanies[i].MotherCompanyID] = struct{}{}
			locks.companiesMapMutex.Unlock()
			motherCompanies, err := motherCompanies[i].MotherCompany(
				qm.Load(models.CompanyRels.MotherCompanyCompanyToCompanies, companyToCompanyFilter),
				qm.Load(models.CompanyRels.CompanyToPeople, companyToPeopleFilter),
				qm.Load(models.CompanyRels.BadCompanies),
			).All(ctx, DB)
			if err != nil {
				writeError(err, response, locks)
				continue
			}
			aggregatedMotherCompanies = append(aggregatedMotherCompanies, motherCompanies...)
		}
	}

	traverseThroughCompanies(ctx, &aggregatedMotherCompanies, response, locks, wg)
	var err error
	companiesToAdd := make(interfaces.Companies, len(aggregatedMotherCompanies))
	for i := range aggregatedMotherCompanies {
		companiesToAdd[i], err = convertCompany(ctx, aggregatedMotherCompanies[i])
		if err != nil {
			writeError(err, response, locks)
		}
	}

	locks.companiesMutex.Lock()
	response.Companies = append(response.Companies, companiesToAdd...)
	locks.companiesMutex.Unlock()
}

// Get the owners of the company
func getOwners(ctx context.Context, companies *models.CompanySlice, response *interfaces.ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()
	if companies == nil || len(*companies) == 0 {
		return
	}

	for _, company := range *companies {
		persons, err := company.CompanyToPeople(qm.Load(models.CompanyToPersonRels.Person), companyToPeopleFilter).All(ctx, DB)
		if err != nil {
			writeError(err, response, locks)
			continue
		}
		searchForBadPersons(ctx, persons, response, locks)
	}
}

// Search for badpersons and add persons to the response
func searchForBadPersons(ctx context.Context, companyToPeople models.CompanyToPersonSlice, response *interfaces.ValidationResponse, locks *validationLocks) {
	if len(companyToPeople) == 0 {
		return
	}
	peopleResponse := make(interfaces.People, len(companyToPeople))
	for i := range companyToPeople {

		locks.peopleMapMutex.RLock()
		_, parsedAlready := locks.personMap[companyToPeople[i].PersonID]
		locks.peopleMapMutex.RUnlock()
		if parsedAlready {
			continue
		}
		locks.peopleMapMutex.Lock()
		locks.personMap[companyToPeople[i].PersonID] = struct{}{}
		locks.peopleMapMutex.Unlock()
		person, err := companyToPeople[i].Person().One(ctx, DB)
		if err != nil {
			writeError(err, response, locks)
		}
		peopleResponse[i], err = convertPerson(ctx, person)
		if err != nil {
			writeError(err, response, locks)
		}
	}

	locks.peopleMutex.Lock()
	response.People = append(response.People, peopleResponse...)
	locks.peopleMutex.Unlock()
}

func writeError(err error, response *interfaces.ValidationResponse, locks *validationLocks) {
	if !errors.Is(err, sql.ErrNoRows) {
		locks.errorsMutex.Lock()
		response.Errors = append(response.Errors, err)
		locks.errorsMutex.Unlock()
	}
}
