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

func testBadPersonsAllNames(t *testing.T) {
	t.Parallel()

	query := BadPersonsAllNames()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBadPersonsAllNamesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
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

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsAllNamesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BadPersonsAllNames().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsAllNamesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadPersonsAllNameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsAllNamesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BadPersonsAllNameExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BadPersonsAllName exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BadPersonsAllNameExists to return true, but got false.")
	}
}

func testBadPersonsAllNamesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	badPersonsAllNameFound, err := FindBadPersonsAllName(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if badPersonsAllNameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBadPersonsAllNamesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BadPersonsAllNames().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBadPersonsAllNamesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BadPersonsAllNames().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBadPersonsAllNamesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	badPersonsAllNameOne := &BadPersonsAllName{}
	badPersonsAllNameTwo := &BadPersonsAllName{}
	if err = randomize.Struct(seed, badPersonsAllNameOne, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}
	if err = randomize.Struct(seed, badPersonsAllNameTwo, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badPersonsAllNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badPersonsAllNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadPersonsAllNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBadPersonsAllNamesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	badPersonsAllNameOne := &BadPersonsAllName{}
	badPersonsAllNameTwo := &BadPersonsAllName{}
	if err = randomize.Struct(seed, badPersonsAllNameOne, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}
	if err = randomize.Struct(seed, badPersonsAllNameTwo, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badPersonsAllNameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badPersonsAllNameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBadPersonsAllNamesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadPersonsAllNamesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(badPersonsAllNameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadPersonsAllNameToOneBadPersonUsingBadPerson(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BadPersonsAllName
	var foreign BadPerson

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, badPersonDBTypes, false, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BadPersonID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BadPerson().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BadPersonsAllNameSlice{&local}
	if err = local.L.LoadBadPerson(ctx, tx, false, (*[]*BadPersonsAllName)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BadPerson == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BadPerson = nil
	if err = local.L.LoadBadPerson(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BadPerson == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBadPersonsAllNameToOneSetOpBadPersonUsingBadPerson(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPersonsAllName
	var b, c BadPerson

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonsAllNameDBTypes, false, strmangle.SetComplement(badPersonsAllNamePrimaryKeyColumns, badPersonsAllNameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BadPerson{&b, &c} {
		err = a.SetBadPerson(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BadPerson != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BadPersonsAllNames[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BadPersonID != x.ID {
			t.Error("foreign key was wrong value", a.BadPersonID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BadPersonID))
		reflect.Indirect(reflect.ValueOf(&a.BadPersonID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BadPersonID != x.ID {
			t.Error("foreign key was wrong value", a.BadPersonID, x.ID)
		}
	}
}

func testBadPersonsAllNamesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
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

func testBadPersonsAllNamesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadPersonsAllNameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBadPersonsAllNamesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadPersonsAllNames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	badPersonsAllNameDBTypes = map[string]string{`ID`: `uuid`, `FullName`: `text`, `BadPersonID`: `uuid`, `UpdatedAt`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `NameVector`: `tsvector`}
	_                        = bytes.MinRead
)

func testBadPersonsAllNamesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(badPersonsAllNamePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(badPersonsAllNameAllColumns) == len(badPersonsAllNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBadPersonsAllNamesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(badPersonsAllNameAllColumns) == len(badPersonsAllNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAllName{}
	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badPersonsAllNameDBTypes, true, badPersonsAllNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(badPersonsAllNameAllColumns, badPersonsAllNamePrimaryKeyColumns) {
		fields = badPersonsAllNameAllColumns
	} else {
		fields = strmangle.SetComplement(
			badPersonsAllNameAllColumns,
			badPersonsAllNamePrimaryKeyColumns,
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

	slice := BadPersonsAllNameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBadPersonsAllNamesUpsert(t *testing.T) {
	t.Parallel()

	if len(badPersonsAllNameAllColumns) == len(badPersonsAllNamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BadPersonsAllName{}
	if err = randomize.Struct(seed, &o, badPersonsAllNameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadPersonsAllName: %s", err)
	}

	count, err := BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, badPersonsAllNameDBTypes, false, badPersonsAllNamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAllName struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadPersonsAllName: %s", err)
	}

	count, err = BadPersonsAllNames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
