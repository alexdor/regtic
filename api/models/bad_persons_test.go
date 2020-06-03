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

func testBadPersons(t *testing.T) {
	t.Parallel()

	query := BadPersons()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBadPersonsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
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

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BadPersons().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadPersonSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBadPersonsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BadPersonExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BadPerson exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BadPersonExists to return true, but got false.")
	}
}

func testBadPersonsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	badPersonFound, err := FindBadPerson(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if badPersonFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBadPersonsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BadPersons().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBadPersonsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BadPersons().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBadPersonsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	badPersonOne := &BadPerson{}
	badPersonTwo := &BadPerson{}
	if err = randomize.Struct(seed, badPersonOne, badPersonDBTypes, false, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}
	if err = randomize.Struct(seed, badPersonTwo, badPersonDBTypes, false, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badPersonOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badPersonTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadPersons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBadPersonsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	badPersonOne := &BadPerson{}
	badPersonTwo := &BadPerson{}
	if err = randomize.Struct(seed, badPersonOne, badPersonDBTypes, false, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}
	if err = randomize.Struct(seed, badPersonTwo, badPersonDBTypes, false, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = badPersonOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = badPersonTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBadPersonsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadPersonsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(badPersonColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBadPersonToManyPersons(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, personDBTypes, false, personColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, personDBTypes, false, personColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"bad_person_to_person\" (\"bad_person_id\", \"person_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"bad_person_to_person\" (\"bad_person_id\", \"person_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Persons().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BadPersonSlice{&a}
	if err = a.L.LoadPersons(ctx, tx, false, (*[]*BadPerson)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Persons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Persons = nil
	if err = a.L.LoadPersons(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Persons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBadPersonToManyBadPersonsAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c BadPersonsAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, badPersonsAddressDBTypes, false, badPersonsAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, badPersonsAddressDBTypes, false, badPersonsAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BadPersonID, a.ID)
	queries.Assign(&c.BadPersonID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BadPersonsAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BadPersonID, b.BadPersonID) {
			bFound = true
		}
		if queries.Equal(v.BadPersonID, c.BadPersonID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BadPersonSlice{&a}
	if err = a.L.LoadBadPersonsAddresses(ctx, tx, false, (*[]*BadPerson)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadPersonsAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BadPersonsAddresses = nil
	if err = a.L.LoadBadPersonsAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadPersonsAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBadPersonToManyBadPersonsAllNames(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c BadPersonsAllName

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, badPersonsAllNameDBTypes, false, badPersonsAllNameColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BadPersonID = a.ID
	c.BadPersonID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BadPersonsAllNames().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BadPersonID == b.BadPersonID {
			bFound = true
		}
		if v.BadPersonID == c.BadPersonID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BadPersonSlice{&a}
	if err = a.L.LoadBadPersonsAllNames(ctx, tx, false, (*[]*BadPerson)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadPersonsAllNames); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BadPersonsAllNames = nil
	if err = a.L.LoadBadPersonsAllNames(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BadPersonsAllNames); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBadPersonToManyAddOpPersons(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Person{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Person{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPersons(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.BadPersons[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.BadPersons[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Persons[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Persons[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Persons().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBadPersonToManySetOpPersons(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Person{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
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

	err = a.SetPersons(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Persons().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPersons(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Persons().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.BadPersons) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.BadPersons) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.BadPersons[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.BadPersons[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Persons[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Persons[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBadPersonToManyRemoveOpPersons(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Person{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddPersons(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Persons().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePersons(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Persons().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.BadPersons) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.BadPersons) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.BadPersons[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BadPersons[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Persons) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Persons[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Persons[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBadPersonToManyAddOpBadPersonsAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e BadPersonsAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadPersonsAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badPersonsAddressDBTypes, false, strmangle.SetComplement(badPersonsAddressPrimaryKeyColumns, badPersonsAddressColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BadPersonsAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBadPersonsAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.BadPersonID) {
			t.Error("foreign key was wrong value", a.ID, first.BadPersonID)
		}
		if !queries.Equal(a.ID, second.BadPersonID) {
			t.Error("foreign key was wrong value", a.ID, second.BadPersonID)
		}

		if first.R.BadPerson != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BadPerson != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BadPersonsAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BadPersonsAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BadPersonsAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBadPersonToManySetOpBadPersonsAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e BadPersonsAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadPersonsAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badPersonsAddressDBTypes, false, strmangle.SetComplement(badPersonsAddressPrimaryKeyColumns, badPersonsAddressColumnsWithoutDefault)...); err != nil {
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

	err = a.SetBadPersonsAddresses(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BadPersonsAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBadPersonsAddresses(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BadPersonsAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BadPersonID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BadPersonID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.BadPersonID) {
		t.Error("foreign key was wrong value", a.ID, d.BadPersonID)
	}
	if !queries.Equal(a.ID, e.BadPersonID) {
		t.Error("foreign key was wrong value", a.ID, e.BadPersonID)
	}

	if b.R.BadPerson != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BadPerson != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BadPerson != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BadPerson != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.BadPersonsAddresses[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.BadPersonsAddresses[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBadPersonToManyRemoveOpBadPersonsAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e BadPersonsAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadPersonsAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badPersonsAddressDBTypes, false, strmangle.SetComplement(badPersonsAddressPrimaryKeyColumns, badPersonsAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBadPersonsAddresses(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BadPersonsAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBadPersonsAddresses(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BadPersonsAddresses().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BadPersonID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BadPersonID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.BadPerson != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BadPerson != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BadPerson != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.BadPerson != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.BadPersonsAddresses) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.BadPersonsAddresses[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.BadPersonsAddresses[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBadPersonToManyAddOpBadPersonsAllNames(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BadPerson
	var b, c, d, e BadPersonsAllName

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, badPersonDBTypes, false, strmangle.SetComplement(badPersonPrimaryKeyColumns, badPersonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BadPersonsAllName{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, badPersonsAllNameDBTypes, false, strmangle.SetComplement(badPersonsAllNamePrimaryKeyColumns, badPersonsAllNameColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BadPersonsAllName{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBadPersonsAllNames(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BadPersonID {
			t.Error("foreign key was wrong value", a.ID, first.BadPersonID)
		}
		if a.ID != second.BadPersonID {
			t.Error("foreign key was wrong value", a.ID, second.BadPersonID)
		}

		if first.R.BadPerson != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BadPerson != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BadPersonsAllNames[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BadPersonsAllNames[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BadPersonsAllNames().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBadPersonsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
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

func testBadPersonsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BadPersonSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBadPersonsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BadPersons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	badPersonDBTypes = map[string]string{`ID`: `uuid`, `FullName`: `text`, `Type`: `enum.bad_person_type('PEP','SANCTION')`, `Source`: `text`, `UpdatedAt`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `NameVector`: `tsvector`, `CitizenshipCountryCode`: `ARRAYcharacter varying`}
	_                = bytes.MinRead
)

func testBadPersonsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(badPersonPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(badPersonAllColumns) == len(badPersonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBadPersonsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(badPersonAllColumns) == len(badPersonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BadPerson{}
	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, badPersonDBTypes, true, badPersonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(badPersonAllColumns, badPersonPrimaryKeyColumns) {
		fields = badPersonAllColumns
	} else {
		fields = strmangle.SetComplement(
			badPersonAllColumns,
			badPersonPrimaryKeyColumns,
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

	slice := BadPersonSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBadPersonsUpsert(t *testing.T) {
	t.Parallel()

	if len(badPersonAllColumns) == len(badPersonPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BadPerson{}
	if err = randomize.Struct(seed, &o, badPersonDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadPerson: %s", err)
	}

	count, err := BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, badPersonDBTypes, false, badPersonPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BadPerson struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BadPerson: %s", err)
	}

	count, err = BadPersons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
