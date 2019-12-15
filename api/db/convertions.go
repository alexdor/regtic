package db

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/alexdor/regtic/api/interfaces"
	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func DBCompanyToCompanyJson(company *models.Company) interfaces.CompanyJson {
	return interfaces.CompanyJson{
		ID:           company.ID,
		Address:      company.Address,
		Vat:          company.Vat,
		StartingDate: company.StartingDate,
		CountryCode:  company.CountryCode,
		Name:         company.Name,
		Status:       company.Status,
		StatusNotes:  company.StatusNotes,
		Type:         company.Type,
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
