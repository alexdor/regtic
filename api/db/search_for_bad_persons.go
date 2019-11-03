package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type ErrorSlice struct {
	sync.Mutex
	items []error
}

type StringSlice struct {
	sync.RWMutex
	items []string
}

func GetBadPersons(ctx context.Context, vat string) ([]string, []error) {
	company, err := models.Companies(models.CompanyWhere.Vat.EQ(vat), qm.Select(models.CompanyColumns.ID)).One(ctx, DB)
	if err != nil {
		return []string{}, []error{err}
	}
	var badPersons StringSlice
	var errors ErrorSlice
	ids := []string{company.ID}
	var wg sync.WaitGroup
	fmt.Println(sqlQuery)
	wg.Add(1)
	searchForBadPersons(ctx, ids, &badPersons, &errors, &wg)
	wg.Wait()
	return badPersons.items, nil
}

func searchForBadPersons(ctx context.Context, ids []string, badPersons *StringSlice, errors *ErrorSlice, wg *sync.WaitGroup) {
	wg.Add(2)
	go getDaughterCompanies(ctx, ids, badPersons, errors, wg)
	go getOwners(ctx, ids, badPersons, errors, wg)
	wg.Done()
}

func getDaughterCompanies(ctx context.Context, ids []string, badPersons *StringSlice, errors *ErrorSlice, group *sync.WaitGroup) {
	defer group.Done()
	companies, err := models.Companies(models.CompanyWhere.ID.IN(ids), qm.Load(models.CompanyRels.DaughterCompanyCompanies)).All(ctx, DB)
	fmt.Println(companies)
	panic("shit hit the fan mate")
	// DaughterCompanyCompanies(qm.WhereIn("mother_company in ?", ids), qm.Select("daughter_company")).All(ctx, DB)
	writeError(err, errors)
	if len(companies) == 0 {
		return
	}
	companyIDS := make([]string, len(companies))
	for i := range companies {
		companyIDS[i] = companies[i].ID
	}
	group.Add(1)
	go searchForBadPersons(ctx, companyIDS, badPersons, errors, group)

}

//TODO: Add aliases to this query
var sqlQuery = "select b." + models.BadPersonColumns.FullName + " from " + models.TableNames.Companies +
	" companies as c where company.id in $1 inner join " + models.TableNames.CompanyToPerson + " as cp on c." + models.CompanyColumns.ID + " = cp.person_id inner join " +
	models.TableNames.Persons + " as p on p." + models.PersonColumns.ID + " = cp.person_id inner join on " +
	models.TableNames.BadPersons + " as b on ( concat(p." + models.PersonColumns.FirstName + ", p." + models.PersonColumns.LastName + ") = " + models.BadPersonColumns.FullName +
	" or concat(p." + models.PersonColumns.LastName + ", p." + models.PersonColumns.FirstName + ") = " + models.BadPersonColumns.FullName + " ))"

func getOwners(ctx context.Context, ids []string, badPersons *StringSlice, errors *ErrorSlice, wg *sync.WaitGroup) {
	defer wg.Done()
	//companies , err := models.Companies(qm.WhereIn("company in ?",ids,qm.Load(models.CompanyRels.Persons)))..All(ctx,db.DB)
	var badPersonList models.BadPersonSlice

	err := models.NewQuery(qm.SQL(sqlQuery, ids)).Bind(ctx, DB, &badPersonList)
	writeError(err, errors)
	if badPersonList != nil {
		badPersons.Lock()
		for i := range badPersonList {
			if badPersonList[i].FullName.Valid {
				badPersons.items = append(badPersons.items, badPersonList[i].FullName.String)
			}
		}
	}

}

func writeError(err error, errors *ErrorSlice) {
	if err != nil {
		errors.Lock()
		errors.items = append(errors.items, err)
		errors.Unlock()
	}
}
