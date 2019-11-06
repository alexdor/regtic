// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Person is an object representing the database table.
type Person struct {
	ID          string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirstName   null.String `boil:"first_name" json:"first_name,omitempty" toml:"first_name" yaml:"first_name,omitempty"`
	LastName    null.String `boil:"last_name" json:"last_name,omitempty" toml:"last_name" yaml:"last_name,omitempty"`
	CountryCode null.String `boil:"country_code" json:"country_code,omitempty" toml:"country_code" yaml:"country_code,omitempty"`
	UpdatedAt   time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	NameVector  null.String `boil:"name_vector" json:"name_vector,omitempty" toml:"name_vector" yaml:"name_vector,omitempty"`

	R *personR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L personL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PersonColumns = struct {
	ID          string
	FirstName   string
	LastName    string
	CountryCode string
	UpdatedAt   string
	CreatedAt   string
	NameVector  string
}{
	ID:          "id",
	FirstName:   "first_name",
	LastName:    "last_name",
	CountryCode: "country_code",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
	NameVector:  "name_vector",
}

// Generated where

var PersonWhere = struct {
	ID          whereHelperstring
	FirstName   whereHelpernull_String
	LastName    whereHelpernull_String
	CountryCode whereHelpernull_String
	UpdatedAt   whereHelpertime_Time
	CreatedAt   whereHelpertime_Time
	NameVector  whereHelpernull_String
}{
	ID:          whereHelperstring{field: "\"persons\".\"id\""},
	FirstName:   whereHelpernull_String{field: "\"persons\".\"first_name\""},
	LastName:    whereHelpernull_String{field: "\"persons\".\"last_name\""},
	CountryCode: whereHelpernull_String{field: "\"persons\".\"country_code\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"persons\".\"updated_at\""},
	CreatedAt:   whereHelpertime_Time{field: "\"persons\".\"created_at\""},
	NameVector:  whereHelpernull_String{field: "\"persons\".\"name_vector\""},
}

// PersonRels is where relationship names are stored.
var PersonRels = struct {
	Companies string
}{
	Companies: "Companies",
}

// personR is where relationships are stored.
type personR struct {
	Companies CompanySlice
}

// NewStruct creates a new relationship struct
func (*personR) NewStruct() *personR {
	return &personR{}
}

// personL is where Load methods for each relationship are stored.
type personL struct{}

var (
	personAllColumns            = []string{"id", "first_name", "last_name", "country_code", "updated_at", "created_at", "name_vector"}
	personColumnsWithoutDefault = []string{"first_name", "last_name", "country_code", "name_vector"}
	personColumnsWithDefault    = []string{"id", "updated_at", "created_at"}
	personPrimaryKeyColumns     = []string{"id"}
)

type (
	// PersonSlice is an alias for a slice of pointers to Person.
	// This should generally be used opposed to []Person.
	PersonSlice []*Person

	personQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	personType                 = reflect.TypeOf(&Person{})
	personMapping              = queries.MakeStructMapping(personType)
	personPrimaryKeyMapping, _ = queries.BindMapping(personType, personMapping, personPrimaryKeyColumns)
	personInsertCacheMut       sync.RWMutex
	personInsertCache          = make(map[string]insertCache)
	personUpdateCacheMut       sync.RWMutex
	personUpdateCache          = make(map[string]updateCache)
	personUpsertCacheMut       sync.RWMutex
	personUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single person record from the query.
func (q personQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Person, error) {
	o := &Person{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for persons")
	}

	return o, nil
}

// All returns all Person records from the query.
func (q personQuery) All(ctx context.Context, exec boil.ContextExecutor) (PersonSlice, error) {
	var o []*Person

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Person slice")
	}

	return o, nil
}

// Count returns the count of all Person records in the query.
func (q personQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count persons rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q personQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if persons exists")
	}

	return count > 0, nil
}

// Companies retrieves all the company's Companies with an executor.
func (o *Person) Companies(mods ...qm.QueryMod) companyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"company_to_person\" on \"companies\".\"id\" = \"company_to_person\".\"company_id\""),
		qm.Where("\"company_to_person\".\"person_id\"=?", o.ID),
	)

	query := Companies(queryMods...)
	queries.SetFrom(query.Query, "\"companies\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"companies\".*"})
	}

	return query
}

// LoadCompanies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (personL) LoadCompanies(ctx context.Context, e boil.ContextExecutor, singular bool, maybePerson interface{}, mods queries.Applicator) error {
	var slice []*Person
	var object *Person

	if singular {
		object = maybePerson.(*Person)
	} else {
		slice = *maybePerson.(*[]*Person)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &personR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &personR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"companies\".*, \"a\".\"person_id\""),
		qm.From("\"companies\""),
		qm.InnerJoin("\"company_to_person\" as \"a\" on \"companies\".\"id\" = \"a\".\"company_id\""),
		qm.WhereIn("\"a\".\"person_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load companies")
	}

	var resultSlice []*Company

	var localJoinCols []string
	for results.Next() {
		one := new(Company)
		var localJoinCol string

		err = results.Scan(&one.ID, &one.Address, &one.Vat, &one.StartingDate, &one.CountryCode, &one.UpdatedAt, &one.CreatedAt, &one.Name, &one.Status, &one.StatusNotes, &one.NameVector, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for companies")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice companies")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on companies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for companies")
	}

	if singular {
		object.R.Companies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &companyR{}
			}
			foreign.R.Persons = append(foreign.R.Persons, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Companies = append(local.R.Companies, foreign)
				if foreign.R == nil {
					foreign.R = &companyR{}
				}
				foreign.R.Persons = append(foreign.R.Persons, local)
				break
			}
		}
	}

	return nil
}

// AddCompanies adds the given related objects to the existing relationships
// of the person, optionally inserting them as new records.
// Appends related to o.R.Companies.
// Sets related.R.Persons appropriately.
func (o *Person) AddCompanies(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Company) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"company_to_person\" (\"person_id\", \"company_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &personR{
			Companies: related,
		}
	} else {
		o.R.Companies = append(o.R.Companies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &companyR{
				Persons: PersonSlice{o},
			}
		} else {
			rel.R.Persons = append(rel.R.Persons, o)
		}
	}
	return nil
}

// SetCompanies removes all previously related items of the
// person replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Persons's Companies accordingly.
// Replaces o.R.Companies with related.
// Sets related.R.Persons's Companies accordingly.
func (o *Person) SetCompanies(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Company) error {
	query := "delete from \"company_to_person\" where \"person_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeCompaniesFromPersonsSlice(o, related)
	if o.R != nil {
		o.R.Companies = nil
	}
	return o.AddCompanies(ctx, exec, insert, related...)
}

// RemoveCompanies relationships from objects passed in.
// Removes related items from R.Companies (uses pointer comparison, removal does not keep order)
// Sets related.R.Persons.
func (o *Person) RemoveCompanies(ctx context.Context, exec boil.ContextExecutor, related ...*Company) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"company_to_person\" where \"person_id\" = $1 and \"company_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeCompaniesFromPersonsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Companies {
			if rel != ri {
				continue
			}

			ln := len(o.R.Companies)
			if ln > 1 && i < ln-1 {
				o.R.Companies[i] = o.R.Companies[ln-1]
			}
			o.R.Companies = o.R.Companies[:ln-1]
			break
		}
	}

	return nil
}

func removeCompaniesFromPersonsSlice(o *Person, related []*Company) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Persons {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Persons)
			if ln > 1 && i < ln-1 {
				rel.R.Persons[i] = rel.R.Persons[ln-1]
			}
			rel.R.Persons = rel.R.Persons[:ln-1]
			break
		}
	}
}

// Persons retrieves all the records using an executor.
func Persons(mods ...qm.QueryMod) personQuery {
	mods = append(mods, qm.From("\"persons\""))
	return personQuery{NewQuery(mods...)}
}

// FindPerson retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPerson(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Person, error) {
	personObj := &Person{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"persons\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, personObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from persons")
	}

	return personObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Person) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no persons provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(personColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	personInsertCacheMut.RLock()
	cache, cached := personInsertCache[key]
	personInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			personAllColumns,
			personColumnsWithDefault,
			personColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(personType, personMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(personType, personMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"persons\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"persons\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into persons")
	}

	if !cached {
		personInsertCacheMut.Lock()
		personInsertCache[key] = cache
		personInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Person.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Person) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	personUpdateCacheMut.RLock()
	cache, cached := personUpdateCache[key]
	personUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			personAllColumns,
			personPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update persons, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"persons\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, personPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(personType, personMapping, append(wl, personPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update persons row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for persons")
	}

	if !cached {
		personUpdateCacheMut.Lock()
		personUpdateCache[key] = cache
		personUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q personQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for persons")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for persons")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PersonSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), personPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"persons\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, personPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in person slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all person")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Person) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no persons provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(personColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	personUpsertCacheMut.RLock()
	cache, cached := personUpsertCache[key]
	personUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			personAllColumns,
			personColumnsWithDefault,
			personColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			personAllColumns,
			personPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert persons, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(personPrimaryKeyColumns))
			copy(conflict, personPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"persons\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(personType, personMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(personType, personMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert persons")
	}

	if !cached {
		personUpsertCacheMut.Lock()
		personUpsertCache[key] = cache
		personUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Person record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Person) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Person provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), personPrimaryKeyMapping)
	sql := "DELETE FROM \"persons\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from persons")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for persons")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q personQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no personQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from persons")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for persons")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PersonSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), personPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"persons\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, personPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from person slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for persons")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Person) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPerson(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PersonSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PersonSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), personPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"persons\".* FROM \"persons\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, personPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PersonSlice")
	}

	*o = slice

	return nil
}

// PersonExists checks if the Person row exists.
func PersonExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"persons\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if persons exists")
	}

	return exists, nil
}
