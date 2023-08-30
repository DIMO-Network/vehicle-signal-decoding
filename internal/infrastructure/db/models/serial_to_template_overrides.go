// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// SerialToTemplateOverride is an object representing the database table.
type SerialToTemplateOverride struct {
	Serial       string    `boil:"serial" json:"serial" toml:"serial" yaml:"serial"`
	TemplateName string    `boil:"template_name" json:"template_name" toml:"template_name" yaml:"template_name"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *serialToTemplateOverrideR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serialToTemplateOverrideL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SerialToTemplateOverrideColumns = struct {
	Serial       string
	TemplateName string
	CreatedAt    string
	UpdatedAt    string
}{
	Serial:       "serial",
	TemplateName: "template_name",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var SerialToTemplateOverrideTableColumns = struct {
	Serial       string
	TemplateName string
	CreatedAt    string
	UpdatedAt    string
}{
	Serial:       "serial_to_template_overrides.serial",
	TemplateName: "serial_to_template_overrides.template_name",
	CreatedAt:    "serial_to_template_overrides.created_at",
	UpdatedAt:    "serial_to_template_overrides.updated_at",
}

// Generated where

var SerialToTemplateOverrideWhere = struct {
	Serial       whereHelperstring
	TemplateName whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	Serial:       whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\".\"serial\""},
	TemplateName: whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\".\"template_name\""},
	CreatedAt:    whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\".\"updated_at\""},
}

// SerialToTemplateOverrideRels is where relationship names are stored.
var SerialToTemplateOverrideRels = struct {
}{}

// serialToTemplateOverrideR is where relationships are stored.
type serialToTemplateOverrideR struct {
}

// NewStruct creates a new relationship struct
func (*serialToTemplateOverrideR) NewStruct() *serialToTemplateOverrideR {
	return &serialToTemplateOverrideR{}
}

// serialToTemplateOverrideL is where Load methods for each relationship are stored.
type serialToTemplateOverrideL struct{}

var (
	serialToTemplateOverrideAllColumns            = []string{"serial", "template_name", "created_at", "updated_at"}
	serialToTemplateOverrideColumnsWithoutDefault = []string{"serial", "template_name"}
	serialToTemplateOverrideColumnsWithDefault    = []string{"created_at", "updated_at"}
	serialToTemplateOverridePrimaryKeyColumns     = []string{"serial", "template_name"}
	serialToTemplateOverrideGeneratedColumns      = []string{}
)

type (
	// SerialToTemplateOverrideSlice is an alias for a slice of pointers to SerialToTemplateOverride.
	// This should almost always be used instead of []SerialToTemplateOverride.
	SerialToTemplateOverrideSlice []*SerialToTemplateOverride
	// SerialToTemplateOverrideHook is the signature for custom SerialToTemplateOverride hook methods
	SerialToTemplateOverrideHook func(context.Context, boil.ContextExecutor, *SerialToTemplateOverride) error

	serialToTemplateOverrideQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serialToTemplateOverrideType                 = reflect.TypeOf(&SerialToTemplateOverride{})
	serialToTemplateOverrideMapping              = queries.MakeStructMapping(serialToTemplateOverrideType)
	serialToTemplateOverridePrimaryKeyMapping, _ = queries.BindMapping(serialToTemplateOverrideType, serialToTemplateOverrideMapping, serialToTemplateOverridePrimaryKeyColumns)
	serialToTemplateOverrideInsertCacheMut       sync.RWMutex
	serialToTemplateOverrideInsertCache          = make(map[string]insertCache)
	serialToTemplateOverrideUpdateCacheMut       sync.RWMutex
	serialToTemplateOverrideUpdateCache          = make(map[string]updateCache)
	serialToTemplateOverrideUpsertCacheMut       sync.RWMutex
	serialToTemplateOverrideUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var serialToTemplateOverrideAfterSelectHooks []SerialToTemplateOverrideHook

var serialToTemplateOverrideBeforeInsertHooks []SerialToTemplateOverrideHook
var serialToTemplateOverrideAfterInsertHooks []SerialToTemplateOverrideHook

var serialToTemplateOverrideBeforeUpdateHooks []SerialToTemplateOverrideHook
var serialToTemplateOverrideAfterUpdateHooks []SerialToTemplateOverrideHook

var serialToTemplateOverrideBeforeDeleteHooks []SerialToTemplateOverrideHook
var serialToTemplateOverrideAfterDeleteHooks []SerialToTemplateOverrideHook

var serialToTemplateOverrideBeforeUpsertHooks []SerialToTemplateOverrideHook
var serialToTemplateOverrideAfterUpsertHooks []SerialToTemplateOverrideHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SerialToTemplateOverride) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SerialToTemplateOverride) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SerialToTemplateOverride) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SerialToTemplateOverride) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SerialToTemplateOverride) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SerialToTemplateOverride) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SerialToTemplateOverride) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SerialToTemplateOverride) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SerialToTemplateOverride) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serialToTemplateOverrideAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSerialToTemplateOverrideHook registers your hook function for all future operations.
func AddSerialToTemplateOverrideHook(hookPoint boil.HookPoint, serialToTemplateOverrideHook SerialToTemplateOverrideHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		serialToTemplateOverrideAfterSelectHooks = append(serialToTemplateOverrideAfterSelectHooks, serialToTemplateOverrideHook)
	case boil.BeforeInsertHook:
		serialToTemplateOverrideBeforeInsertHooks = append(serialToTemplateOverrideBeforeInsertHooks, serialToTemplateOverrideHook)
	case boil.AfterInsertHook:
		serialToTemplateOverrideAfterInsertHooks = append(serialToTemplateOverrideAfterInsertHooks, serialToTemplateOverrideHook)
	case boil.BeforeUpdateHook:
		serialToTemplateOverrideBeforeUpdateHooks = append(serialToTemplateOverrideBeforeUpdateHooks, serialToTemplateOverrideHook)
	case boil.AfterUpdateHook:
		serialToTemplateOverrideAfterUpdateHooks = append(serialToTemplateOverrideAfterUpdateHooks, serialToTemplateOverrideHook)
	case boil.BeforeDeleteHook:
		serialToTemplateOverrideBeforeDeleteHooks = append(serialToTemplateOverrideBeforeDeleteHooks, serialToTemplateOverrideHook)
	case boil.AfterDeleteHook:
		serialToTemplateOverrideAfterDeleteHooks = append(serialToTemplateOverrideAfterDeleteHooks, serialToTemplateOverrideHook)
	case boil.BeforeUpsertHook:
		serialToTemplateOverrideBeforeUpsertHooks = append(serialToTemplateOverrideBeforeUpsertHooks, serialToTemplateOverrideHook)
	case boil.AfterUpsertHook:
		serialToTemplateOverrideAfterUpsertHooks = append(serialToTemplateOverrideAfterUpsertHooks, serialToTemplateOverrideHook)
	}
}

// One returns a single serialToTemplateOverride record from the query.
func (q serialToTemplateOverrideQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SerialToTemplateOverride, error) {
	o := &SerialToTemplateOverride{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for serial_to_template_overrides")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SerialToTemplateOverride records from the query.
func (q serialToTemplateOverrideQuery) All(ctx context.Context, exec boil.ContextExecutor) (SerialToTemplateOverrideSlice, error) {
	var o []*SerialToTemplateOverride

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SerialToTemplateOverride slice")
	}

	if len(serialToTemplateOverrideAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SerialToTemplateOverride records in the query.
func (q serialToTemplateOverrideQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count serial_to_template_overrides rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serialToTemplateOverrideQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if serial_to_template_overrides exists")
	}

	return count > 0, nil
}

// SerialToTemplateOverrides retrieves all the records using an executor.
func SerialToTemplateOverrides(mods ...qm.QueryMod) serialToTemplateOverrideQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\".*"})
	}

	return serialToTemplateOverrideQuery{q}
}

// FindSerialToTemplateOverride retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSerialToTemplateOverride(ctx context.Context, exec boil.ContextExecutor, serial string, templateName string, selectCols ...string) (*SerialToTemplateOverride, error) {
	serialToTemplateOverrideObj := &SerialToTemplateOverride{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" where \"serial\"=$1 AND \"template_name\"=$2", sel,
	)

	q := queries.Raw(query, serial, templateName)

	err := q.Bind(ctx, exec, serialToTemplateOverrideObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from serial_to_template_overrides")
	}

	if err = serialToTemplateOverrideObj.doAfterSelectHooks(ctx, exec); err != nil {
		return serialToTemplateOverrideObj, err
	}

	return serialToTemplateOverrideObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SerialToTemplateOverride) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no serial_to_template_overrides provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serialToTemplateOverrideColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serialToTemplateOverrideInsertCacheMut.RLock()
	cache, cached := serialToTemplateOverrideInsertCache[key]
	serialToTemplateOverrideInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serialToTemplateOverrideAllColumns,
			serialToTemplateOverrideColumnsWithDefault,
			serialToTemplateOverrideColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serialToTemplateOverrideType, serialToTemplateOverrideMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serialToTemplateOverrideType, serialToTemplateOverrideMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into serial_to_template_overrides")
	}

	if !cached {
		serialToTemplateOverrideInsertCacheMut.Lock()
		serialToTemplateOverrideInsertCache[key] = cache
		serialToTemplateOverrideInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SerialToTemplateOverride.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SerialToTemplateOverride) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	serialToTemplateOverrideUpdateCacheMut.RLock()
	cache, cached := serialToTemplateOverrideUpdateCache[key]
	serialToTemplateOverrideUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serialToTemplateOverrideAllColumns,
			serialToTemplateOverridePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update serial_to_template_overrides, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serialToTemplateOverridePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serialToTemplateOverrideType, serialToTemplateOverrideMapping, append(wl, serialToTemplateOverridePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update serial_to_template_overrides row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for serial_to_template_overrides")
	}

	if !cached {
		serialToTemplateOverrideUpdateCacheMut.Lock()
		serialToTemplateOverrideUpdateCache[key] = cache
		serialToTemplateOverrideUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q serialToTemplateOverrideQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for serial_to_template_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for serial_to_template_overrides")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SerialToTemplateOverrideSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serialToTemplateOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serialToTemplateOverridePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serialToTemplateOverride slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serialToTemplateOverride")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SerialToTemplateOverride) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no serial_to_template_overrides provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(serialToTemplateOverrideColumnsWithDefault, o)

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

	serialToTemplateOverrideUpsertCacheMut.RLock()
	cache, cached := serialToTemplateOverrideUpsertCache[key]
	serialToTemplateOverrideUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serialToTemplateOverrideAllColumns,
			serialToTemplateOverrideColumnsWithDefault,
			serialToTemplateOverrideColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			serialToTemplateOverrideAllColumns,
			serialToTemplateOverridePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert serial_to_template_overrides, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serialToTemplateOverridePrimaryKeyColumns))
			copy(conflict, serialToTemplateOverridePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"serial_to_template_overrides\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serialToTemplateOverrideType, serialToTemplateOverrideMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serialToTemplateOverrideType, serialToTemplateOverrideMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert serial_to_template_overrides")
	}

	if !cached {
		serialToTemplateOverrideUpsertCacheMut.Lock()
		serialToTemplateOverrideUpsertCache[key] = cache
		serialToTemplateOverrideUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SerialToTemplateOverride record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SerialToTemplateOverride) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SerialToTemplateOverride provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serialToTemplateOverridePrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" WHERE \"serial\"=$1 AND \"template_name\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from serial_to_template_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for serial_to_template_overrides")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serialToTemplateOverrideQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serialToTemplateOverrideQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serial_to_template_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for serial_to_template_overrides")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SerialToTemplateOverrideSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(serialToTemplateOverrideBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serialToTemplateOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serialToTemplateOverridePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serialToTemplateOverride slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for serial_to_template_overrides")
	}

	if len(serialToTemplateOverrideAfterDeleteHooks) != 0 {
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
func (o *SerialToTemplateOverride) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSerialToTemplateOverride(ctx, exec, o.Serial, o.TemplateName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SerialToTemplateOverrideSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SerialToTemplateOverrideSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serialToTemplateOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\".* FROM \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serialToTemplateOverridePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SerialToTemplateOverrideSlice")
	}

	*o = slice

	return nil
}

// SerialToTemplateOverrideExists checks if the SerialToTemplateOverride row exists.
func SerialToTemplateOverrideExists(ctx context.Context, exec boil.ContextExecutor, serial string, templateName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"serial_to_template_overrides\" where \"serial\"=$1 AND \"template_name\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, serial, templateName)
	}
	row := exec.QueryRowContext(ctx, sql, serial, templateName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if serial_to_template_overrides exists")
	}

	return exists, nil
}

// Exists checks if the SerialToTemplateOverride row exists.
func (o *SerialToTemplateOverride) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SerialToTemplateOverrideExists(ctx, exec, o.Serial, o.TemplateName)
}