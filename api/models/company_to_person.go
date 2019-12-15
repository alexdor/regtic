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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// CompanyToPerson is an object representing the database table.
type CompanyToPerson struct {
	CompanyID    string            `boil:"company_id" json:"companyID" toml:"companyID" yaml:"companyID"`
	PersonID     string            `boil:"person_id" json:"personID" toml:"personID" yaml:"personID"`
	Relations    types.StringArray `boil:"relations" json:"relations" toml:"relations" yaml:"relations"`
	VotingRights float32           `boil:"voting_rights" json:"votingRights" toml:"votingRights" yaml:"votingRights"`
	Ownership    float32           `boil:"ownership" json:"ownership" toml:"ownership" yaml:"ownership"`

	R *companyToPersonR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L companyToPersonL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompanyToPersonColumns = struct {
	CompanyID    string
	PersonID     string
	Relations    string
	VotingRights string
	Ownership    string
}{
	CompanyID:    "company_id",
	PersonID:     "person_id",
	Relations:    "relations",
	VotingRights: "voting_rights",
	Ownership:    "ownership",
}

// Generated where

var CompanyToPersonWhere = struct {
	CompanyID    whereHelperstring
	PersonID     whereHelperstring
	Relations    whereHelpertypes_StringArray
	VotingRights whereHelperfloat32
	Ownership    whereHelperfloat32
}{
	CompanyID:    whereHelperstring{field: "\"company_to_person\".\"company_id\""},
	PersonID:     whereHelperstring{field: "\"company_to_person\".\"person_id\""},
	Relations:    whereHelpertypes_StringArray{field: "\"company_to_person\".\"relations\""},
	VotingRights: whereHelperfloat32{field: "\"company_to_person\".\"voting_rights\""},
	Ownership:    whereHelperfloat32{field: "\"company_to_person\".\"ownership\""},
}

// CompanyToPersonRels is where relationship names are stored.
var CompanyToPersonRels = struct {
	Company string
	Person  string
}{
	Company: "Company",
	Person:  "Person",
}

// companyToPersonR is where relationships are stored.
type companyToPersonR struct {
	Company *Company
	Person  *Person
}

// NewStruct creates a new relationship struct
func (*companyToPersonR) NewStruct() *companyToPersonR {
	return &companyToPersonR{}
}

// companyToPersonL is where Load methods for each relationship are stored.
type companyToPersonL struct{}

var (
	companyToPersonAllColumns            = []string{"company_id", "person_id", "relations", "voting_rights", "ownership"}
	companyToPersonColumnsWithoutDefault = []string{"company_id", "person_id"}
	companyToPersonColumnsWithDefault    = []string{"relations", "voting_rights", "ownership"}
	companyToPersonPrimaryKeyColumns     = []string{"company_id", "person_id"}
)

type (
	// CompanyToPersonSlice is an alias for a slice of pointers to CompanyToPerson.
	// This should generally be used opposed to []CompanyToPerson.
	CompanyToPersonSlice []*CompanyToPerson

	companyToPersonQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	companyToPersonType                 = reflect.TypeOf(&CompanyToPerson{})
	companyToPersonMapping              = queries.MakeStructMapping(companyToPersonType)
	companyToPersonPrimaryKeyMapping, _ = queries.BindMapping(companyToPersonType, companyToPersonMapping, companyToPersonPrimaryKeyColumns)
	companyToPersonInsertCacheMut       sync.RWMutex
	companyToPersonInsertCache          = make(map[string]insertCache)
	companyToPersonUpdateCacheMut       sync.RWMutex
	companyToPersonUpdateCache          = make(map[string]updateCache)
	companyToPersonUpsertCacheMut       sync.RWMutex
	companyToPersonUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single companyToPerson record from the query.
func (q companyToPersonQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CompanyToPerson, error) {
	o := &CompanyToPerson{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for company_to_person")
	}

	return o, nil
}

// All returns all CompanyToPerson records from the query.
func (q companyToPersonQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompanyToPersonSlice, error) {
	var o []*CompanyToPerson

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CompanyToPerson slice")
	}

	return o, nil
}

// Count returns the count of all CompanyToPerson records in the query.
func (q companyToPersonQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count company_to_person rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q companyToPersonQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if company_to_person exists")
	}

	return count > 0, nil
}

// Company pointed to by the foreign key.
func (o *CompanyToPerson) Company(mods ...qm.QueryMod) companyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CompanyID),
	}

	queryMods = append(queryMods, mods...)

	query := Companies(queryMods...)
	queries.SetFrom(query.Query, "\"companies\"")

	return query
}

// Person pointed to by the foreign key.
func (o *CompanyToPerson) Person(mods ...qm.QueryMod) personQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PersonID),
	}

	queryMods = append(queryMods, mods...)

	query := Persons(queryMods...)
	queries.SetFrom(query.Query, "\"persons\"")

	return query
}

// LoadCompany allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (companyToPersonL) LoadCompany(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCompanyToPerson interface{}, mods queries.Applicator) error {
	var slice []*CompanyToPerson
	var object *CompanyToPerson

	if singular {
		object = maybeCompanyToPerson.(*CompanyToPerson)
	} else {
		slice = *maybeCompanyToPerson.(*[]*CompanyToPerson)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &companyToPersonR{}
		}
		args = append(args, object.CompanyID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &companyToPersonR{}
			}

			for _, a := range args {
				if a == obj.CompanyID {
					continue Outer
				}
			}

			args = append(args, obj.CompanyID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`companies`), qm.WhereIn(`companies.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Company")
	}

	var resultSlice []*Company
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Company")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for companies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for companies")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Company = foreign
		if foreign.R == nil {
			foreign.R = &companyR{}
		}
		foreign.R.CompanyToPeople = append(foreign.R.CompanyToPeople, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CompanyID == foreign.ID {
				local.R.Company = foreign
				if foreign.R == nil {
					foreign.R = &companyR{}
				}
				foreign.R.CompanyToPeople = append(foreign.R.CompanyToPeople, local)
				break
			}
		}
	}

	return nil
}

// LoadPerson allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (companyToPersonL) LoadPerson(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCompanyToPerson interface{}, mods queries.Applicator) error {
	var slice []*CompanyToPerson
	var object *CompanyToPerson

	if singular {
		object = maybeCompanyToPerson.(*CompanyToPerson)
	} else {
		slice = *maybeCompanyToPerson.(*[]*CompanyToPerson)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &companyToPersonR{}
		}
		args = append(args, object.PersonID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &companyToPersonR{}
			}

			for _, a := range args {
				if a == obj.PersonID {
					continue Outer
				}
			}

			args = append(args, obj.PersonID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`persons`), qm.WhereIn(`persons.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Person")
	}

	var resultSlice []*Person
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Person")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for persons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for persons")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Person = foreign
		if foreign.R == nil {
			foreign.R = &personR{}
		}
		foreign.R.CompanyToPeople = append(foreign.R.CompanyToPeople, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PersonID == foreign.ID {
				local.R.Person = foreign
				if foreign.R == nil {
					foreign.R = &personR{}
				}
				foreign.R.CompanyToPeople = append(foreign.R.CompanyToPeople, local)
				break
			}
		}
	}

	return nil
}

// SetCompany of the companyToPerson to the related item.
// Sets o.R.Company to related.
// Adds o to related.R.CompanyToPeople.
func (o *CompanyToPerson) SetCompany(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Company) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"company_to_person\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"company_id"}),
		strmangle.WhereClause("\"", "\"", 2, companyToPersonPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.CompanyID, o.PersonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CompanyID = related.ID
	if o.R == nil {
		o.R = &companyToPersonR{
			Company: related,
		}
	} else {
		o.R.Company = related
	}

	if related.R == nil {
		related.R = &companyR{
			CompanyToPeople: CompanyToPersonSlice{o},
		}
	} else {
		related.R.CompanyToPeople = append(related.R.CompanyToPeople, o)
	}

	return nil
}

// SetPerson of the companyToPerson to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.CompanyToPeople.
func (o *CompanyToPerson) SetPerson(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Person) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"company_to_person\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"person_id"}),
		strmangle.WhereClause("\"", "\"", 2, companyToPersonPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.CompanyID, o.PersonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PersonID = related.ID
	if o.R == nil {
		o.R = &companyToPersonR{
			Person: related,
		}
	} else {
		o.R.Person = related
	}

	if related.R == nil {
		related.R = &personR{
			CompanyToPeople: CompanyToPersonSlice{o},
		}
	} else {
		related.R.CompanyToPeople = append(related.R.CompanyToPeople, o)
	}

	return nil
}

// CompanyToPeople retrieves all the records using an executor.
func CompanyToPeople(mods ...qm.QueryMod) companyToPersonQuery {
	mods = append(mods, qm.From("\"company_to_person\""))
	return companyToPersonQuery{NewQuery(mods...)}
}

// FindCompanyToPerson retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompanyToPerson(ctx context.Context, exec boil.ContextExecutor, companyID string, personID string, selectCols ...string) (*CompanyToPerson, error) {
	companyToPersonObj := &CompanyToPerson{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"company_to_person\" where \"company_id\"=$1 AND \"person_id\"=$2", sel,
	)

	q := queries.Raw(query, companyID, personID)

	err := q.Bind(ctx, exec, companyToPersonObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from company_to_person")
	}

	return companyToPersonObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompanyToPerson) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company_to_person provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(companyToPersonColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	companyToPersonInsertCacheMut.RLock()
	cache, cached := companyToPersonInsertCache[key]
	companyToPersonInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			companyToPersonAllColumns,
			companyToPersonColumnsWithDefault,
			companyToPersonColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(companyToPersonType, companyToPersonMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(companyToPersonType, companyToPersonMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"company_to_person\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"company_to_person\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into company_to_person")
	}

	if !cached {
		companyToPersonInsertCacheMut.Lock()
		companyToPersonInsertCache[key] = cache
		companyToPersonInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CompanyToPerson.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompanyToPerson) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	companyToPersonUpdateCacheMut.RLock()
	cache, cached := companyToPersonUpdateCache[key]
	companyToPersonUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			companyToPersonAllColumns,
			companyToPersonPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update company_to_person, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"company_to_person\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, companyToPersonPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(companyToPersonType, companyToPersonMapping, append(wl, companyToPersonPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update company_to_person row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for company_to_person")
	}

	if !cached {
		companyToPersonUpdateCacheMut.Lock()
		companyToPersonUpdateCache[key] = cache
		companyToPersonUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q companyToPersonQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for company_to_person")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for company_to_person")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompanyToPersonSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyToPersonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"company_to_person\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, companyToPersonPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in companyToPerson slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all companyToPerson")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompanyToPerson) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company_to_person provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(companyToPersonColumnsWithDefault, o)

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

	companyToPersonUpsertCacheMut.RLock()
	cache, cached := companyToPersonUpsertCache[key]
	companyToPersonUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			companyToPersonAllColumns,
			companyToPersonColumnsWithDefault,
			companyToPersonColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			companyToPersonAllColumns,
			companyToPersonPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert company_to_person, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(companyToPersonPrimaryKeyColumns))
			copy(conflict, companyToPersonPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"company_to_person\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(companyToPersonType, companyToPersonMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(companyToPersonType, companyToPersonMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert company_to_person")
	}

	if !cached {
		companyToPersonUpsertCacheMut.Lock()
		companyToPersonUpsertCache[key] = cache
		companyToPersonUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CompanyToPerson record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompanyToPerson) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompanyToPerson provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), companyToPersonPrimaryKeyMapping)
	sql := "DELETE FROM \"company_to_person\" WHERE \"company_id\"=$1 AND \"person_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from company_to_person")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for company_to_person")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q companyToPersonQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no companyToPersonQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from company_to_person")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company_to_person")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompanyToPersonSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyToPersonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"company_to_person\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, companyToPersonPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from companyToPerson slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company_to_person")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CompanyToPerson) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompanyToPerson(ctx, exec, o.CompanyID, o.PersonID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompanyToPersonSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompanyToPersonSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyToPersonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"company_to_person\".* FROM \"company_to_person\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, companyToPersonPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompanyToPersonSlice")
	}

	*o = slice

	return nil
}

// CompanyToPersonExists checks if the CompanyToPerson row exists.
func CompanyToPersonExists(ctx context.Context, exec boil.ContextExecutor, companyID string, personID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"company_to_person\" where \"company_id\"=$1 AND \"person_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, companyID, personID)
	}

	row := exec.QueryRowContext(ctx, sql, companyID, personID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if company_to_person exists")
	}

	return exists, nil
}