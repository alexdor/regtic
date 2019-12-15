package interfaces

import "github.com/alexdor/regtic/api/models"

func DBCompanyToCompanyJson(company *models.Company) CompanyJson {
	return CompanyJson{
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
