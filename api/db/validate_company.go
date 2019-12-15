package db

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"sync"

	"github.com/alexdor/regtic/api/interfaces"
	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/null"
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
		qm.Load(models.CompanyRels.MotherCompanyCompanyToCompanies),
		qm.Load(models.CompanyRels.CompanyToPeople),
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
	traverseThroughTheCompany(ctx, &models.CompanySlice{company}, response, locks, wg)
	response.Info = interfaces.DBCompanyToCompanyJson(company)
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

func traverseThroughTheCompany(ctx context.Context, companies *models.CompanySlice, response *interfaces.ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	wg.Add(2)
	go getMotherCompanies(ctx, companies, response, locks, wg)
	go getOwners(ctx, companies, response, locks, wg)
}

func getMotherCompanies(ctx context.Context, companies *models.CompanySlice, response *interfaces.ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()

	if companies == nil || len(*companies) == 0 {
		return
	}
	aggregatedMotherCompanies := models.CompanySlice{}

	for _, company := range *companies {

		motherCompanies, err := company.MotherCompanyCompanyToCompanies(qm.Load(models.CompanyToCompanyRels.MotherCompany)).All(ctx, DB)
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
				qm.Load(models.CompanyRels.MotherCompanyCompanyToCompanies),
				qm.Load(models.CompanyRels.CompanyToPeople),
				qm.Load(models.CompanyRels.BadCompanies),
			).All(ctx, DB)
			if err != nil {
				writeError(err, response, locks)
				continue
			}
			aggregatedMotherCompanies = append(aggregatedMotherCompanies, motherCompanies...)
		}
	}

	traverseThroughTheCompany(ctx, &aggregatedMotherCompanies, response, locks, wg)
	var err error
	companiesToAdd := make(interfaces.Companies, len(aggregatedMotherCompanies))
	for i := range aggregatedMotherCompanies {
		companiesToAdd[i], err = convertCompany(ctx, aggregatedMotherCompanies[i]) //TODO
		if err != nil {
			writeError(err, response, locks)
		}
	}

	locks.companiesMutex.Lock()
	response.Companies = append(response.Companies, companiesToAdd...)
	locks.companiesMutex.Unlock()
}

func getOwners(ctx context.Context, companies *models.CompanySlice, response *interfaces.ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()
	if companies == nil || len(*companies) == 0 {
		return
	}

	for _, company := range *companies {
		persons, err := company.CompanyToPeople(qm.Load(models.CompanyToPersonRels.Person)).All(ctx, DB)
		if err != nil {
			writeError(err, response, locks)
			continue
		}
		searchForBadPersons(ctx, persons, response, locks)
	}
}

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

func convertPerson(ctx context.Context, person *models.Person) (interfaces.Person, error) {
	jsonPerson := interfaces.Person{
		ID:          person.ID,
		CountryCode: person.CountryCode,
		Name:        person.FullName,
		EntityType:  interfaces.PERSON,
		CheckStatus: interfaces.OK,
	}
	badPerson, err := models.BadPersons(
		qm.Select(badPersonSelect),
		qm.Where(badPersonWhereClause, person.FullName),
		qm.OrderBy(badPersonOrderBy),
	).One(ctx, DB)

	noResults := errors.Is(err, sql.ErrNoRows)
	if err != nil && !noResults {
		return jsonPerson, err
	}

	if !noResults {
		jsonPerson.CheckStatus = interfaces.BadTypeToStatusMapping[strings.ToLower(badPerson.Type)]
		jsonPerson.Source = badPerson.Source
		jsonPerson.BadType = null.String{String: badPerson.Type, Valid: true}
	}

	return jsonPerson, nil
}

func convertCompany(ctx context.Context, company *models.Company) (interfaces.Company, error) {
	jsonCompany := interfaces.Company{
		ID:           company.ID,
		Address:      company.Address,
		CountryCode:  company.CountryCode,
		Name:         company.Name,
		Vat:          company.Vat,
		StartingDate: company.StartingDate,
		Status:       company.Status,
		StatusNotes:  company.StatusNotes,
		Type:         company.Type,
		EntityType:   interfaces.COMPANY,
		CheckStatus:  interfaces.OK,
	}
	companyOwners, err := company.MotherCompanyCompanyToCompanies().All(ctx, DB)
	noResults := errors.Is(err, sql.ErrNoRows)
	if err != nil && !noResults {
		return jsonCompany, err
	}
	peopleOwners, err := company.CompanyToPeople().All(ctx, DB)
	noResults = errors.Is(err, sql.ErrNoRows)
	if err != nil && !noResults {
		return jsonCompany, err
	}
	peopleOwnersLength := len(peopleOwners)
	ownedBy := make([]interfaces.Relationship, peopleOwnersLength+len(companyOwners))
	for i := range peopleOwners {
		ownedBy[i] = interfaces.Relationship{
			EntityID:     peopleOwners[i].PersonID,
			Relation:     peopleOwners[i].Relations,
			Ownership:    peopleOwners[i].Ownership,
			VotingRights: peopleOwners[i].VotingRights,
			EntityType:   interfaces.PERSON,
		}
	}
	for i := range companyOwners {
		ownedBy[peopleOwnersLength-1+i] = interfaces.Relationship{
			EntityID:     companyOwners[i].MotherCompanyID,
			Relation:     companyOwners[i].Relations,
			Ownership:    companyOwners[i].Ownership,
			VotingRights: companyOwners[i].VotingRights,
			EntityType:   interfaces.COMPANY,
		}
	}
	jsonCompany.OwnedBy = ownedBy

	badCompany, err := company.BadCompanies(
		qm.OrderBy(badCompanyOrderBy),
	).One(ctx, DB)

	noResults = errors.Is(err, sql.ErrNoRows)
	if err != nil && !noResults {
		return jsonCompany, err
	}

	if !noResults {
		jsonCompany.Source = null.String{String: badCompany.Source, Valid: true}
		jsonCompany.BadType = null.String{String: badCompany.Type, Valid: true}
		jsonCompany.CheckStatus = interfaces.BadTypeToStatusMapping[strings.ToLower(badCompany.Type)]
	}

	return jsonCompany, nil
}
