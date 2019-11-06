package db

import (
	"context"
	"sync"

	"github.com/alexdor/regtic/api/interfaces"
	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type validationLocks struct {
	companies sync.RWMutex
	people    struct {
		bad     sync.RWMutex
		warning sync.RWMutex
		good    sync.RWMutex
	}
	errors sync.Mutex
}

type ValidationResponse struct {
	Info      models.Company      `json:"info"`
	Companies models.CompanySlice `json:"companies"`
	People    struct {
		Bad     []interfaces.BadPersonJson `json:"bad"`
		Warning []interfaces.BadPersonJson `json:"warning"`
		Good    []interfaces.PersonJson    `json:"good"`
	} `json:"people"`
	Errors []error `json:"errors"`
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
	searchForBadPersons(ctx, &models.CompanySlice{company}, &response, locks, wg)

	response.Info = *company

	wg.Wait()

	unique(response)
	return &response, nil
}

func unique(response ValidationResponse) {
	//TODO
	return
}

func searchForBadPersons(ctx context.Context, companies *models.CompanySlice, response *ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	wg.Add(2)
	go getDaughterCompanies(ctx, companies, response, locks, wg)
	go getOwners(ctx, companies, response, locks, wg)
	wg.Done()
}

func getDaughterCompanies(ctx context.Context, companies *models.CompanySlice, response *ValidationResponse, locks *validationLocks, group *sync.WaitGroup) {
	defer group.Done()

	if companies == nil || len(*companies) == 0 {
		return
	}
	motherCompanies := models.CompanySlice{}
	for _, company := range *companies {
		mCompanies, err := company.MotherCompanyCompanies(qm.Load(models.CompanyRels.Persons), qm.Load(models.CompanyRels.MotherCompanyCompanies)).All(ctx, DB)
		motherCompanies = append(motherCompanies, mCompanies...)
		writeError(err, response, locks)
	}
	group.Add(1)
	searchForBadPersons(ctx, &motherCompanies, response, locks, group)

	locks.companies.Lock()
	response.Companies = append(response.Companies, motherCompanies...)
	locks.companies.Unlock()
}

//TODO: Add aliases to this query
var sqlQuery = "select b." + models.BadPersonColumns.FullName + " from " + models.TableNames.Companies +
	" companies as c where company.id in $1 inner join " + models.TableNames.CompanyToPerson + " as cp on c." + models.CompanyColumns.ID + " = cp.person_id inner join " +
	models.TableNames.Persons + " as p on p." + models.PersonColumns.ID + " = cp.person_id inner join on " +
	models.TableNames.BadPersons + " as b on ( concat(p." + models.PersonColumns.FirstName + ", p." + models.PersonColumns.LastName + ") = " + models.BadPersonColumns.FullName +
	" or concat(p." + models.PersonColumns.LastName + ", p." + models.PersonColumns.FirstName + ") = " + models.BadPersonColumns.FullName + " ))"

func getOwners(ctx context.Context, companies *models.CompanySlice, response *ValidationResponse, locks *validationLocks, wg *sync.WaitGroup) {
	defer wg.Done()
	//companies , err := models.Companies(qm.WhereIn("company in ?",ids,qm.Load(models.CompanyRels.Persons)))..All(ctx,db.DB)
	//var badPersonList models.BadPersonSlice
	//
	//err := models.NewQuery(qm.SQL(sqlQuery, ids)).Bind(ctx, DB, &badPersonList)
	//writeError(err, response, locks)
	//if badPersonList != nil {
	//badPersons.Lock()
	//for i := range badPersonList {
	//	if badPersonList[i].FullName.Valid {
	//		badPersons.items = append(badPersons.items, badPersonList[i].FullName.String)
	//	}
	//}
	//}

}

func writeError(err error, response *ValidationResponse, locks *validationLocks) {
	if err != nil {
		locks.errors.Lock()
		response.Errors = append(response.Errors, err)
		locks.errors.Unlock()
	}
}
