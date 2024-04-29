// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DBCFile is an object representing the database table.
type DBCFile struct {
	DBCFile      string    `boil:"dbc_file" json:"dbc_file" toml:"dbc_file" yaml:"dbc_file"`
	TemplateName string    `boil:"template_name" json:"template_name" toml:"template_name" yaml:"template_name"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *dbcFileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dbcFileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DBCFileColumns = struct {
	DBCFile      string
	TemplateName string
	CreatedAt    string
	UpdatedAt    string
}{
	DBCFile:      "dbc_file",
	TemplateName: "template_name",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var DBCFileTableColumns = struct {
	DBCFile      string
	TemplateName string
	CreatedAt    string
	UpdatedAt    string
}{
	DBCFile:      "dbc_files.dbc_file",
	TemplateName: "dbc_files.template_name",
	CreatedAt:    "dbc_files.created_at",
	UpdatedAt:    "dbc_files.updated_at",
}

// Generated where

var DBCFileWhere = struct {
	DBCFile      whereHelperstring
	TemplateName whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	DBCFile:      whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"dbc_files\".\"dbc_file\""},
	TemplateName: whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"dbc_files\".\"template_name\""},
	CreatedAt:    whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"dbc_files\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"dbc_files\".\"updated_at\""},
}

// DBCFileRels is where relationship names are stored.
var DBCFileRels = struct {
	TemplateNameTemplate string
}{
	TemplateNameTemplate: "TemplateNameTemplate",
}

// dbcFileR is where relationships are stored.
type dbcFileR struct {
	TemplateNameTemplate *Template `boil:"TemplateNameTemplate" json:"TemplateNameTemplate" toml:"TemplateNameTemplate" yaml:"TemplateNameTemplate"`
}

// NewStruct creates a new relationship struct
func (*dbcFileR) NewStruct() *dbcFileR {
	return &dbcFileR{}
}

func (r *dbcFileR) GetTemplateNameTemplate() *Template {
	if r == nil {
		return nil
	}
	return r.TemplateNameTemplate
}

// dbcFileL is where Load methods for each relationship are stored.
type dbcFileL struct{}

var (
	dbcFileAllColumns            = []string{"dbc_file", "template_name", "created_at", "updated_at"}
	dbcFileColumnsWithoutDefault = []string{"dbc_file", "template_name"}
	dbcFileColumnsWithDefault    = []string{"created_at", "updated_at"}
	dbcFilePrimaryKeyColumns     = []string{"template_name"}
	dbcFileGeneratedColumns      = []string{}
)

type (
	// DBCFileSlice is an alias for a slice of pointers to DBCFile.
	// This should almost always be used instead of []DBCFile.
	DBCFileSlice []*DBCFile
	// DBCFileHook is the signature for custom DBCFile hook methods
	DBCFileHook func(context.Context, boil.ContextExecutor, *DBCFile) error

	dbcFileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dbcFileType                 = reflect.TypeOf(&DBCFile{})
	dbcFileMapping              = queries.MakeStructMapping(dbcFileType)
	dbcFilePrimaryKeyMapping, _ = queries.BindMapping(dbcFileType, dbcFileMapping, dbcFilePrimaryKeyColumns)
	dbcFileInsertCacheMut       sync.RWMutex
	dbcFileInsertCache          = make(map[string]insertCache)
	dbcFileUpdateCacheMut       sync.RWMutex
	dbcFileUpdateCache          = make(map[string]updateCache)
	dbcFileUpsertCacheMut       sync.RWMutex
	dbcFileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dbcFileAfterSelectMu sync.Mutex
var dbcFileAfterSelectHooks []DBCFileHook

var dbcFileBeforeInsertMu sync.Mutex
var dbcFileBeforeInsertHooks []DBCFileHook
var dbcFileAfterInsertMu sync.Mutex
var dbcFileAfterInsertHooks []DBCFileHook

var dbcFileBeforeUpdateMu sync.Mutex
var dbcFileBeforeUpdateHooks []DBCFileHook
var dbcFileAfterUpdateMu sync.Mutex
var dbcFileAfterUpdateHooks []DBCFileHook

var dbcFileBeforeDeleteMu sync.Mutex
var dbcFileBeforeDeleteHooks []DBCFileHook
var dbcFileAfterDeleteMu sync.Mutex
var dbcFileAfterDeleteHooks []DBCFileHook

var dbcFileBeforeUpsertMu sync.Mutex
var dbcFileBeforeUpsertHooks []DBCFileHook
var dbcFileAfterUpsertMu sync.Mutex
var dbcFileAfterUpsertHooks []DBCFileHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DBCFile) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DBCFile) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DBCFile) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DBCFile) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DBCFile) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DBCFile) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DBCFile) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DBCFile) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DBCFile) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dbcFileAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDBCFileHook registers your hook function for all future operations.
func AddDBCFileHook(hookPoint boil.HookPoint, dbcFileHook DBCFileHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		dbcFileAfterSelectMu.Lock()
		dbcFileAfterSelectHooks = append(dbcFileAfterSelectHooks, dbcFileHook)
		dbcFileAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		dbcFileBeforeInsertMu.Lock()
		dbcFileBeforeInsertHooks = append(dbcFileBeforeInsertHooks, dbcFileHook)
		dbcFileBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		dbcFileAfterInsertMu.Lock()
		dbcFileAfterInsertHooks = append(dbcFileAfterInsertHooks, dbcFileHook)
		dbcFileAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		dbcFileBeforeUpdateMu.Lock()
		dbcFileBeforeUpdateHooks = append(dbcFileBeforeUpdateHooks, dbcFileHook)
		dbcFileBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		dbcFileAfterUpdateMu.Lock()
		dbcFileAfterUpdateHooks = append(dbcFileAfterUpdateHooks, dbcFileHook)
		dbcFileAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		dbcFileBeforeDeleteMu.Lock()
		dbcFileBeforeDeleteHooks = append(dbcFileBeforeDeleteHooks, dbcFileHook)
		dbcFileBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		dbcFileAfterDeleteMu.Lock()
		dbcFileAfterDeleteHooks = append(dbcFileAfterDeleteHooks, dbcFileHook)
		dbcFileAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		dbcFileBeforeUpsertMu.Lock()
		dbcFileBeforeUpsertHooks = append(dbcFileBeforeUpsertHooks, dbcFileHook)
		dbcFileBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		dbcFileAfterUpsertMu.Lock()
		dbcFileAfterUpsertHooks = append(dbcFileAfterUpsertHooks, dbcFileHook)
		dbcFileAfterUpsertMu.Unlock()
	}
}

// One returns a single dbcFile record from the query.
func (q dbcFileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DBCFile, error) {
	o := &DBCFile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dbc_files")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DBCFile records from the query.
func (q dbcFileQuery) All(ctx context.Context, exec boil.ContextExecutor) (DBCFileSlice, error) {
	var o []*DBCFile

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DBCFile slice")
	}

	if len(dbcFileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DBCFile records in the query.
func (q dbcFileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dbc_files rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dbcFileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dbc_files exists")
	}

	return count > 0, nil
}

// TemplateNameTemplate pointed to by the foreign key.
func (o *DBCFile) TemplateNameTemplate(mods ...qm.QueryMod) templateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"template_name\" = ?", o.TemplateName),
	}

	queryMods = append(queryMods, mods...)

	return Templates(queryMods...)
}

// LoadTemplateNameTemplate allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dbcFileL) LoadTemplateNameTemplate(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDBCFile interface{}, mods queries.Applicator) error {
	var slice []*DBCFile
	var object *DBCFile

	if singular {
		var ok bool
		object, ok = maybeDBCFile.(*DBCFile)
		if !ok {
			object = new(DBCFile)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDBCFile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDBCFile))
			}
		}
	} else {
		s, ok := maybeDBCFile.(*[]*DBCFile)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDBCFile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDBCFile))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &dbcFileR{}
		}
		args[object.TemplateName] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dbcFileR{}
			}

			args[obj.TemplateName] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`vehicle_signal_decoding_api.templates`),
		qm.WhereIn(`vehicle_signal_decoding_api.templates.template_name in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Template")
	}

	var resultSlice []*Template
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Template")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for templates")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for templates")
	}

	if len(templateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TemplateNameTemplate = foreign
		if foreign.R == nil {
			foreign.R = &templateR{}
		}
		foreign.R.TemplateNameDBCFile = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TemplateName == foreign.TemplateName {
				local.R.TemplateNameTemplate = foreign
				if foreign.R == nil {
					foreign.R = &templateR{}
				}
				foreign.R.TemplateNameDBCFile = local
				break
			}
		}
	}

	return nil
}

// SetTemplateNameTemplate of the dbcFile to the related item.
// Sets o.R.TemplateNameTemplate to related.
// Adds o to related.R.TemplateNameDBCFile.
func (o *DBCFile) SetTemplateNameTemplate(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Template) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"vehicle_signal_decoding_api\".\"dbc_files\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"template_name"}),
		strmangle.WhereClause("\"", "\"", 2, dbcFilePrimaryKeyColumns),
	)
	values := []interface{}{related.TemplateName, o.TemplateName}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TemplateName = related.TemplateName
	if o.R == nil {
		o.R = &dbcFileR{
			TemplateNameTemplate: related,
		}
	} else {
		o.R.TemplateNameTemplate = related
	}

	if related.R == nil {
		related.R = &templateR{
			TemplateNameDBCFile: o,
		}
	} else {
		related.R.TemplateNameDBCFile = o
	}

	return nil
}

// DBCFiles retrieves all the records using an executor.
func DBCFiles(mods ...qm.QueryMod) dbcFileQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"dbc_files\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"dbc_files\".*"})
	}

	return dbcFileQuery{q}
}

// FindDBCFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDBCFile(ctx context.Context, exec boil.ContextExecutor, templateName string, selectCols ...string) (*DBCFile, error) {
	dbcFileObj := &DBCFile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"dbc_files\" where \"template_name\"=$1", sel,
	)

	q := queries.Raw(query, templateName)

	err := q.Bind(ctx, exec, dbcFileObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dbc_files")
	}

	if err = dbcFileObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dbcFileObj, err
	}

	return dbcFileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DBCFile) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dbc_files provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbcFileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dbcFileInsertCacheMut.RLock()
	cache, cached := dbcFileInsertCache[key]
	dbcFileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dbcFileAllColumns,
			dbcFileColumnsWithDefault,
			dbcFileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dbcFileType, dbcFileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dbcFileType, dbcFileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"dbc_files\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"dbc_files\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into dbc_files")
	}

	if !cached {
		dbcFileInsertCacheMut.Lock()
		dbcFileInsertCache[key] = cache
		dbcFileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DBCFile.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DBCFile) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dbcFileUpdateCacheMut.RLock()
	cache, cached := dbcFileUpdateCache[key]
	dbcFileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dbcFileAllColumns,
			dbcFilePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update dbc_files, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"dbc_files\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dbcFilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dbcFileType, dbcFileMapping, append(wl, dbcFilePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update dbc_files row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for dbc_files")
	}

	if !cached {
		dbcFileUpdateCacheMut.Lock()
		dbcFileUpdateCache[key] = cache
		dbcFileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dbcFileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for dbc_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for dbc_files")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DBCFileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbcFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"dbc_files\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dbcFilePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dbcFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dbcFile")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DBCFile) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no dbc_files provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbcFileColumnsWithDefault, o)

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

	dbcFileUpsertCacheMut.RLock()
	cache, cached := dbcFileUpsertCache[key]
	dbcFileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			dbcFileAllColumns,
			dbcFileColumnsWithDefault,
			dbcFileColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dbcFileAllColumns,
			dbcFilePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert dbc_files, could not build update column list")
		}

		ret := strmangle.SetComplement(dbcFileAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(dbcFilePrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert dbc_files, could not build conflict column list")
			}

			conflict = make([]string, len(dbcFilePrimaryKeyColumns))
			copy(conflict, dbcFilePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"dbc_files\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(dbcFileType, dbcFileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dbcFileType, dbcFileMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert dbc_files")
	}

	if !cached {
		dbcFileUpsertCacheMut.Lock()
		dbcFileUpsertCache[key] = cache
		dbcFileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DBCFile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DBCFile) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DBCFile provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dbcFilePrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"dbc_files\" WHERE \"template_name\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from dbc_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for dbc_files")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dbcFileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dbcFileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dbc_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dbc_files")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DBCFileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(dbcFileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbcFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"dbc_files\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dbcFilePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dbcFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dbc_files")
	}

	if len(dbcFileAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DBCFile) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDBCFile(ctx, exec, o.TemplateName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DBCFileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DBCFileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbcFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"dbc_files\".* FROM \"vehicle_signal_decoding_api\".\"dbc_files\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dbcFilePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DBCFileSlice")
	}

	*o = slice

	return nil
}

// DBCFileExists checks if the DBCFile row exists.
func DBCFileExists(ctx context.Context, exec boil.ContextExecutor, templateName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"dbc_files\" where \"template_name\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, templateName)
	}
	row := exec.QueryRowContext(ctx, sql, templateName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dbc_files exists")
	}

	return exists, nil
}

// Exists checks if the DBCFile row exists.
func (o *DBCFile) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DBCFileExists(ctx, exec, o.TemplateName)
}
