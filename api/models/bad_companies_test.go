// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBadCompanies(t *testing.T) {
	t.Parallel()

	query := BadCompanies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBadCompaniesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadCompaniesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BadCompanies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadCompaniesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadCompanySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadCompaniesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BadCompanyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BadCompany exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BadCompanyExists to return true, but got false.")
	}
}

func testBadCompaniesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	badCompanyFound, err := FindBadCompany(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if badCompanyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBadCompaniesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BadCompanies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBadCompaniesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BadCompanies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBadCompaniesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	badCompanyOne := &BadCompany{}
	badCompanyTwo := &BadCompany{}
	if err = randomize.Struct(seed, badCompanyOne, badCompanyDBTypes, false, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}
	if err = randomize.Struct(seed, badCompanyTwo, badCompanyDBTypes, false, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badCompanyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badCompanyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadCompanies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBadCompaniesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	badCompanyOne := &BadCompany{}
	badCompanyTwo := &BadCompany{}
	if err = randomize.Struct(seed, badCompanyOne, badCompanyDBTypes, false, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}
	if err = randomize.Struct(seed, badCompanyTwo, badCompanyDBTypes, false, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badCompanyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badCompanyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBadCompaniesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadCompaniesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(badCompanyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadCompanyToManyBadCompaniesAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c BadCompaniesAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, badCompaniesAddressDBTypes, false, badCompaniesAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, badCompaniesAddressDBTypes, false, badCompaniesAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BadCompanyID, a.ID)
	queries.Assign(&c.BadCompanyID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BadCompaniesAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BadCompanyID, b.BadCompanyID) {
			bFound = true
		}
		if queries.Equal(v.BadCompanyID, c.BadCompanyID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BadCompanySlice{&a}
	if err = a.L.LoadBadCompaniesAddresses(ctx, tx, false, (*[]*BadCompany)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadCompaniesAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BadCompaniesAddresses = nil
	if err = a.L.LoadBadCompaniesAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadCompaniesAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBadCompanyToManyBadCompaniesAliases(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c BadCompaniesAlias

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, badCompaniesAliasDBTypes, false, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, badCompaniesAliasDBTypes, false, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BadCompanyID, a.ID)
	queries.Assign(&c.BadCompanyID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BadCompaniesAliases().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BadCompanyID, b.BadCompanyID) {
			bFound = true
		}
		if queries.Equal(v.BadCompanyID, c.BadCompanyID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BadCompanySlice{&a}
	if err = a.L.LoadBadCompaniesAliases(ctx, tx, false, (*[]*BadCompany)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadCompaniesAliases); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BadCompaniesAliases = nil
	if err = a.L.LoadBadCompaniesAliases(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadCompaniesAliases); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBadCompanyToManyAddOpBadCompaniesAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c, d, e BadCompaniesAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadCompaniesAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badCompaniesAddressDBTypes, false, strmangle.SetComplement(badCompaniesAddressPrimaryKeyColumns, badCompaniesAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BadCompaniesAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBadCompaniesAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.BadCompanyID) {
			t.Error("foreign key was wrong value", a.ID, first.BadCompanyID)
		}
		if !queries.Equal(a.ID, second.BadCompanyID) {
			t.Error("foreign key was wrong value", a.ID, second.BadCompanyID)
		}

		if first.R.BadCompany != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BadCompany != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BadCompaniesAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BadCompaniesAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BadCompaniesAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBadCompanyToManySetOpBadCompaniesAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c, d, e BadCompaniesAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadCompaniesAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badCompaniesAddressDBTypes, false, strmangle.SetComplement(badCompaniesAddressPrimaryKeyColumns, badCompaniesAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetBadCompaniesAddresses(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BadCompaniesAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBadCompaniesAddresses(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BadCompaniesAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BadCompanyID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BadCompanyID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.BadCompanyID) {
		t.Error("foreign key was wrong value", a.ID, d.BadCompanyID)
	}
	if !queries.Equal(a.ID, e.BadCompanyID) {
		t.Error("foreign key was wrong value", a.ID, e.BadCompanyID)
	}

	if b.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BadCompany != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BadCompany != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.BadCompaniesAddresses[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.BadCompaniesAddresses[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBadCompanyToManyRemoveOpBadCompaniesAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c, d, e BadCompaniesAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadCompaniesAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badCompaniesAddressDBTypes, false, strmangle.SetComplement(badCompaniesAddressPrimaryKeyColumns, badCompaniesAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBadCompaniesAddresses(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BadCompaniesAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBadCompaniesAddresses(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BadCompaniesAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BadCompanyID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BadCompanyID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BadCompany != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.BadCompany != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.BadCompaniesAddresses) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.BadCompaniesAddresses[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.BadCompaniesAddresses[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBadCompanyToManyAddOpBadCompaniesAliases(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c, d, e BadCompaniesAlias

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadCompaniesAlias{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badCompaniesAliasDBTypes, false, strmangle.SetComplement(badCompaniesAliasPrimaryKeyColumns, badCompaniesAliasColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BadCompaniesAlias{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBadCompaniesAliases(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.BadCompanyID) {
			t.Error("foreign key was wrong value", a.ID, first.BadCompanyID)
		}
		if !queries.Equal(a.ID, second.BadCompanyID) {
			t.Error("foreign key was wrong value", a.ID, second.BadCompanyID)
		}

		if first.R.BadCompany != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BadCompany != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BadCompaniesAliases[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BadCompaniesAliases[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BadCompaniesAliases().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBadCompanyToManySetOpBadCompaniesAliases(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c, d, e BadCompaniesAlias

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadCompaniesAlias{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badCompaniesAliasDBTypes, false, strmangle.SetComplement(badCompaniesAliasPrimaryKeyColumns, badCompaniesAliasColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetBadCompaniesAliases(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBadCompaniesAliases(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BadCompanyID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BadCompanyID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.BadCompanyID) {
		t.Error("foreign key was wrong value", a.ID, d.BadCompanyID)
	}
	if !queries.Equal(a.ID, e.BadCompanyID) {
		t.Error("foreign key was wrong value", a.ID, e.BadCompanyID)
	}

	if b.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BadCompany != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BadCompany != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.BadCompaniesAliases[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.BadCompaniesAliases[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBadCompanyToManyRemoveOpBadCompaniesAliases(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompany
	var b, c, d, e BadCompaniesAlias

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadCompaniesAlias{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badCompaniesAliasDBTypes, false, strmangle.SetComplement(badCompaniesAliasPrimaryKeyColumns, badCompaniesAliasColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBadCompaniesAliases(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBadCompaniesAliases(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BadCompanyID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BadCompanyID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BadCompany != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BadCompany != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.BadCompany != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.BadCompaniesAliases) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.BadCompaniesAliases[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.BadCompaniesAliases[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBadCompaniesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBadCompaniesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadCompanySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBadCompaniesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadCompanies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	badCompanyDBTypes = map[string]string{`ID`: `uuid`, `UpdatedAt`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `Name`: `text`, `NameVector`: `tsvector`, `Address`: `text`, `Source`: `text`, `CitizenshipRegion`: `text`, `Type`: `enum.bad_person_type('PEP','SANCTION')`, `CitizenshipCountryCode`: `ARRAYcharacter varying`}
	_                 = bytes.MinRead
)

func testBadCompaniesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(badCompanyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(badCompanyAllColumns) == len(badCompanyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBadCompaniesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(badCompanyAllColumns) == len(badCompanyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadCompany{}
	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badCompanyDBTypes, true, badCompanyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(badCompanyAllColumns, badCompanyPrimaryKeyColumns) {
		fields = badCompanyAllColumns
	} else {
		fields = strmangle.SetComplement(
			badCompanyAllColumns,
			badCompanyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BadCompanySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBadCompaniesUpsert(t *testing.T) {
	t.Parallel()

	if len(badCompanyAllColumns) == len(badCompanyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BadCompany{}
	if err = randomize.Struct(seed, &o, badCompanyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadCompany: %s", err)
	}

	count, err := BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, badCompanyDBTypes, false, badCompanyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadCompany: %s", err)
	}

	count, err = BadCompanies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}