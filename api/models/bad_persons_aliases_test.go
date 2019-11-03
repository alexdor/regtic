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

func testBadPersonsAliases(t *testing.T) {
	t.Parallel()

	query := BadPersonsAliases()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBadPersonsAliasesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
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

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsAliasesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BadPersonsAliases().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsAliasesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadPersonsAliasSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsAliasesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BadPersonsAliasExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BadPersonsAlias exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BadPersonsAliasExists to return true, but got false.")
	}
}

func testBadPersonsAliasesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	badPersonsAliasFound, err := FindBadPersonsAlias(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if badPersonsAliasFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBadPersonsAliasesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BadPersonsAliases().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBadPersonsAliasesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BadPersonsAliases().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBadPersonsAliasesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	badPersonsAliasOne := &BadPersonsAlias{}
	badPersonsAliasTwo := &BadPersonsAlias{}
	if err = randomize.Struct(seed, badPersonsAliasOne, badPersonsAliasDBTypes, false, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}
	if err = randomize.Struct(seed, badPersonsAliasTwo, badPersonsAliasDBTypes, false, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badPersonsAliasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badPersonsAliasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadPersonsAliases().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBadPersonsAliasesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	badPersonsAliasOne := &BadPersonsAlias{}
	badPersonsAliasTwo := &BadPersonsAlias{}
	if err = randomize.Struct(seed, badPersonsAliasOne, badPersonsAliasDBTypes, false, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}
	if err = randomize.Struct(seed, badPersonsAliasTwo, badPersonsAliasDBTypes, false, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badPersonsAliasOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badPersonsAliasTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func badPersonsAliasBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func badPersonsAliasAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BadPersonsAlias) error {
	*o = BadPersonsAlias{}
	return nil
}

func testBadPersonsAliasesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BadPersonsAlias{}
	o := &BadPersonsAlias{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias object: %s", err)
	}

	AddBadPersonsAliasHook(boil.BeforeInsertHook, badPersonsAliasBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasBeforeInsertHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.AfterInsertHook, badPersonsAliasAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasAfterInsertHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.AfterSelectHook, badPersonsAliasAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasAfterSelectHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.BeforeUpdateHook, badPersonsAliasBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasBeforeUpdateHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.AfterUpdateHook, badPersonsAliasAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasAfterUpdateHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.BeforeDeleteHook, badPersonsAliasBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasBeforeDeleteHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.AfterDeleteHook, badPersonsAliasAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasAfterDeleteHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.BeforeUpsertHook, badPersonsAliasBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasBeforeUpsertHooks = []BadPersonsAliasHook{}

	AddBadPersonsAliasHook(boil.AfterUpsertHook, badPersonsAliasAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	badPersonsAliasAfterUpsertHooks = []BadPersonsAliasHook{}
}

func testBadPersonsAliasesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadPersonsAliasesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(badPersonsAliasColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadPersonsAliasToOneBadPersonUsingBadPerson(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BadPersonsAlias
	var foreign BadPerson

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, badPersonsAliasDBTypes, false, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
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

	slice := BadPersonsAliasSlice{&local}
	if err = local.L.LoadBadPerson(ctx, tx, false, (*[]*BadPersonsAlias)(&slice), nil); err != nil {
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

func testBadPersonsAliasToOneSetOpBadPersonUsingBadPerson(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPersonsAlias
	var b, c BadPerson

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonsAliasDBTypes, false, strmangle.SetComplement(badPersonsAliasPrimaryKeyColumns, badPersonsAliasColumnsWithoutDefault)...); err != nil {
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

		if x.R.BadPersonsAliases[0] != &a {
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

func testBadPersonsAliasesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
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

func testBadPersonsAliasesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadPersonsAliasSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBadPersonsAliasesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadPersonsAliases().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	badPersonsAliasDBTypes = map[string]string{`ID`: `uuid`, `FullName`: `text`, `BadPersonID`: `uuid`, `UpdatedAt`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `NameVector`: `tsvector`}
	_                      = bytes.MinRead
)

func testBadPersonsAliasesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(badPersonsAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(badPersonsAliasAllColumns) == len(badPersonsAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBadPersonsAliasesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(badPersonsAliasAllColumns) == len(badPersonsAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadPersonsAlias{}
	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badPersonsAliasDBTypes, true, badPersonsAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(badPersonsAliasAllColumns, badPersonsAliasPrimaryKeyColumns) {
		fields = badPersonsAliasAllColumns
	} else {
		fields = strmangle.SetComplement(
			badPersonsAliasAllColumns,
			badPersonsAliasPrimaryKeyColumns,
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

	slice := BadPersonsAliasSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBadPersonsAliasesUpsert(t *testing.T) {
	t.Parallel()

	if len(badPersonsAliasAllColumns) == len(badPersonsAliasPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BadPersonsAlias{}
	if err = randomize.Struct(seed, &o, badPersonsAliasDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadPersonsAlias: %s", err)
	}

	count, err := BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, badPersonsAliasDBTypes, false, badPersonsAliasPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPersonsAlias struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadPersonsAlias: %s", err)
	}

	count, err = BadPersonsAliases().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
