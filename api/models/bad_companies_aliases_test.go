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

func testBadCompaniesAliases(t *testing.T) {
	t.Parallel()

	query := BadCompaniesAliases()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBadCompaniesAliasesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
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

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadCompaniesAliasesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BadCompaniesAliases().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadCompaniesAliasesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadCompaniesAliasSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadCompaniesAliasesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BadCompaniesAliasExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BadCompaniesAlias exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BadCompaniesAliasExists to return true, but got false.")
	}
}

func testBadCompaniesAliasesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	badCompaniesAliasFound, err := FindBadCompaniesAlias(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if badCompaniesAliasFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBadCompaniesAliasesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BadCompaniesAliases().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBadCompaniesAliasesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BadCompaniesAliases().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBadCompaniesAliasesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	badCompaniesAliasOne := &BadCompaniesAlias{}
	badCompaniesAliasTwo := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, badCompaniesAliasOne, badCompaniesAliasDBTypes, false, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}
	if err = randomize.Struct(seed, badCompaniesAliasTwo, badCompaniesAliasDBTypes, false, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badCompaniesAliasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badCompaniesAliasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadCompaniesAliases().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBadCompaniesAliasesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	badCompaniesAliasOne := &BadCompaniesAlias{}
	badCompaniesAliasTwo := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, badCompaniesAliasOne, badCompaniesAliasDBTypes, false, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}
	if err = randomize.Struct(seed, badCompaniesAliasTwo, badCompaniesAliasDBTypes, false, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badCompaniesAliasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badCompaniesAliasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBadCompaniesAliasesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadCompaniesAliasesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(badCompaniesAliasColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadCompaniesAliasToOneBadCompanyUsingBadCompany(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BadCompaniesAlias
	var foreign BadCompany

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, badCompanyDBTypes, false, badCompanyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompany struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BadCompanyID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BadCompany().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BadCompaniesAliasSlice{&local}
	if err = local.L.LoadBadCompany(ctx, tx, false, (*[]*BadCompaniesAlias)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BadCompany == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BadCompany = nil
	if err = local.L.LoadBadCompany(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BadCompany == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBadCompaniesAliasToOneSetOpBadCompanyUsingBadCompany(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompaniesAlias
	var b, c BadCompany

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompaniesAliasDBTypes, false, strmangle.SetComplement(badCompaniesAliasPrimaryKeyColumns, badCompaniesAliasColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BadCompany{&b, &c} {
		err = a.SetBadCompany(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BadCompany != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BadCompaniesAliases[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BadCompanyID, x.ID) {
			t.Error("foreign key was wrong value", a.BadCompanyID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BadCompanyID))
		reflect.Indirect(reflect.ValueOf(&a.BadCompanyID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BadCompanyID, x.ID) {
			t.Error("foreign key was wrong value", a.BadCompanyID, x.ID)
		}
	}
}

func testBadCompaniesAliasToOneRemoveOpBadCompanyUsingBadCompany(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadCompaniesAlias
	var b BadCompany

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badCompaniesAliasDBTypes, false, strmangle.SetComplement(badCompaniesAliasPrimaryKeyColumns, badCompaniesAliasColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, badCompanyDBTypes, false, strmangle.SetComplement(badCompanyPrimaryKeyColumns, badCompanyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBadCompany(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBadCompany(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.BadCompany().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.BadCompany != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BadCompanyID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.BadCompaniesAliases) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBadCompaniesAliasesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
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

func testBadCompaniesAliasesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadCompaniesAliasSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBadCompaniesAliasesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadCompaniesAliases().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	badCompaniesAliasDBTypes = map[string]string{`ID`: `uuid`, `UpdatedAt`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `BadCompanyID`: `uuid`, `Name`: `text`, `NameVector`: `tsvector`, `Type`: `enum.bad_person_type('PEP','SANCTION')`}
	_                        = bytes.MinRead
)

func testBadCompaniesAliasesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(badCompaniesAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(badCompaniesAliasAllColumns) == len(badCompaniesAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBadCompaniesAliasesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(badCompaniesAliasAllColumns) == len(badCompaniesAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadCompaniesAlias{}
	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badCompaniesAliasDBTypes, true, badCompaniesAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(badCompaniesAliasAllColumns, badCompaniesAliasPrimaryKeyColumns) {
		fields = badCompaniesAliasAllColumns
	} else {
		fields = strmangle.SetComplement(
			badCompaniesAliasAllColumns,
			badCompaniesAliasPrimaryKeyColumns,
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

	slice := BadCompaniesAliasSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBadCompaniesAliasesUpsert(t *testing.T) {
	t.Parallel()

	if len(badCompaniesAliasAllColumns) == len(badCompaniesAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BadCompaniesAlias{}
	if err = randomize.Struct(seed, &o, badCompaniesAliasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadCompaniesAlias: %s", err)
	}

	count, err := BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, badCompaniesAliasDBTypes, false, badCompaniesAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadCompaniesAlias struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadCompaniesAlias: %s", err)
	}

	count, err = BadCompaniesAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
