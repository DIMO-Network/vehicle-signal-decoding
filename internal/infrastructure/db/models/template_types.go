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

// TemplateType is an object representing the database table.
type TemplateType struct {
	TypeName  string    `boil:"type_name" json:"type_name" toml:"type_name" yaml:"type_name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *templateTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L templateTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TemplateTypeColumns = struct {
	TypeName  string
	CreatedAt string
	UpdatedAt string
}{
	TypeName:  "type_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var TemplateTypeTableColumns = struct {
	TypeName  string
	CreatedAt string
	UpdatedAt string
}{
	TypeName:  "template_types.type_name",
	CreatedAt: "template_types.created_at",
	UpdatedAt: "template_types.updated_at",
}

// Generated where

var TemplateTypeWhere = struct {
	TypeName  whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	TypeName:  whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"template_types\".\"type_name\""},
	CreatedAt: whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"template_types\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"template_types\".\"updated_at\""},
}

// TemplateTypeRels is where relationship names are stored.
var TemplateTypeRels = struct {
}{}

// templateTypeR is where relationships are stored.
type templateTypeR struct {
}

// NewStruct creates a new relationship struct
func (*templateTypeR) NewStruct() *templateTypeR {
	return &templateTypeR{}
}

// templateTypeL is where Load methods for each relationship are stored.
type templateTypeL struct{}

var (
	templateTypeAllColumns            = []string{"type_name", "created_at", "updated_at"}
	templateTypeColumnsWithoutDefault = []string{"type_name"}
	templateTypeColumnsWithDefault    = []string{"created_at", "updated_at"}
	templateTypePrimaryKeyColumns     = []string{"type_name"}
	templateTypeGeneratedColumns      = []string{}
)

type (
	// TemplateTypeSlice is an alias for a slice of pointers to TemplateType.
	// This should almost always be used instead of []TemplateType.
	TemplateTypeSlice []*TemplateType
	// TemplateTypeHook is the signature for custom TemplateType hook methods
	TemplateTypeHook func(context.Context, boil.ContextExecutor, *TemplateType) error

	templateTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	templateTypeType                 = reflect.TypeOf(&TemplateType{})
	templateTypeMapping              = queries.MakeStructMapping(templateTypeType)
	templateTypePrimaryKeyMapping, _ = queries.BindMapping(templateTypeType, templateTypeMapping, templateTypePrimaryKeyColumns)
	templateTypeInsertCacheMut       sync.RWMutex
	templateTypeInsertCache          = make(map[string]insertCache)
	templateTypeUpdateCacheMut       sync.RWMutex
	templateTypeUpdateCache          = make(map[string]updateCache)
	templateTypeUpsertCacheMut       sync.RWMutex
	templateTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var templateTypeAfterSelectHooks []TemplateTypeHook

var templateTypeBeforeInsertHooks []TemplateTypeHook
var templateTypeAfterInsertHooks []TemplateTypeHook

var templateTypeBeforeUpdateHooks []TemplateTypeHook
var templateTypeAfterUpdateHooks []TemplateTypeHook

var templateTypeBeforeDeleteHooks []TemplateTypeHook
var templateTypeAfterDeleteHooks []TemplateTypeHook

var templateTypeBeforeUpsertHooks []TemplateTypeHook
var templateTypeAfterUpsertHooks []TemplateTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TemplateType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TemplateType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TemplateType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TemplateType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TemplateType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TemplateType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TemplateType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TemplateType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TemplateType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTemplateTypeHook registers your hook function for all future operations.
func AddTemplateTypeHook(hookPoint boil.HookPoint, templateTypeHook TemplateTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		templateTypeAfterSelectHooks = append(templateTypeAfterSelectHooks, templateTypeHook)
	case boil.BeforeInsertHook:
		templateTypeBeforeInsertHooks = append(templateTypeBeforeInsertHooks, templateTypeHook)
	case boil.AfterInsertHook:
		templateTypeAfterInsertHooks = append(templateTypeAfterInsertHooks, templateTypeHook)
	case boil.BeforeUpdateHook:
		templateTypeBeforeUpdateHooks = append(templateTypeBeforeUpdateHooks, templateTypeHook)
	case boil.AfterUpdateHook:
		templateTypeAfterUpdateHooks = append(templateTypeAfterUpdateHooks, templateTypeHook)
	case boil.BeforeDeleteHook:
		templateTypeBeforeDeleteHooks = append(templateTypeBeforeDeleteHooks, templateTypeHook)
	case boil.AfterDeleteHook:
		templateTypeAfterDeleteHooks = append(templateTypeAfterDeleteHooks, templateTypeHook)
	case boil.BeforeUpsertHook:
		templateTypeBeforeUpsertHooks = append(templateTypeBeforeUpsertHooks, templateTypeHook)
	case boil.AfterUpsertHook:
		templateTypeAfterUpsertHooks = append(templateTypeAfterUpsertHooks, templateTypeHook)
	}
}

// One returns a single templateType record from the query.
func (q templateTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TemplateType, error) {
	o := &TemplateType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for template_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TemplateType records from the query.
func (q templateTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (TemplateTypeSlice, error) {
	var o []*TemplateType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TemplateType slice")
	}

	if len(templateTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TemplateType records in the query.
func (q templateTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count template_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q templateTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if template_types exists")
	}

	return count > 0, nil
}

// TemplateTypes retrieves all the records using an executor.
func TemplateTypes(mods ...qm.QueryMod) templateTypeQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"template_types\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"template_types\".*"})
	}

	return templateTypeQuery{q}
}

// FindTemplateType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTemplateType(ctx context.Context, exec boil.ContextExecutor, typeName string, selectCols ...string) (*TemplateType, error) {
	templateTypeObj := &TemplateType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"template_types\" where \"type_name\"=$1", sel,
	)

	q := queries.Raw(query, typeName)

	err := q.Bind(ctx, exec, templateTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from template_types")
	}

	if err = templateTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return templateTypeObj, err
	}

	return templateTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TemplateType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no template_types provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(templateTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	templateTypeInsertCacheMut.RLock()
	cache, cached := templateTypeInsertCache[key]
	templateTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			templateTypeAllColumns,
			templateTypeColumnsWithDefault,
			templateTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(templateTypeType, templateTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(templateTypeType, templateTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"template_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"template_types\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into template_types")
	}

	if !cached {
		templateTypeInsertCacheMut.Lock()
		templateTypeInsertCache[key] = cache
		templateTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TemplateType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TemplateType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	templateTypeUpdateCacheMut.RLock()
	cache, cached := templateTypeUpdateCache[key]
	templateTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			templateTypeAllColumns,
			templateTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update template_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"template_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, templateTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(templateTypeType, templateTypeMapping, append(wl, templateTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update template_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for template_types")
	}

	if !cached {
		templateTypeUpdateCacheMut.Lock()
		templateTypeUpdateCache[key] = cache
		templateTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q templateTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for template_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for template_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TemplateTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), templateTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"template_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, templateTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in templateType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all templateType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TemplateType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no template_types provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(templateTypeColumnsWithDefault, o)

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

	templateTypeUpsertCacheMut.RLock()
	cache, cached := templateTypeUpsertCache[key]
	templateTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			templateTypeAllColumns,
			templateTypeColumnsWithDefault,
			templateTypeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			templateTypeAllColumns,
			templateTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert template_types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(templateTypePrimaryKeyColumns))
			copy(conflict, templateTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"template_types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(templateTypeType, templateTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(templateTypeType, templateTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert template_types")
	}

	if !cached {
		templateTypeUpsertCacheMut.Lock()
		templateTypeUpsertCache[key] = cache
		templateTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TemplateType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TemplateType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TemplateType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), templateTypePrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"template_types\" WHERE \"type_name\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from template_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for template_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q templateTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no templateTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from template_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for template_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TemplateTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(templateTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), templateTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"template_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, templateTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from templateType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for template_types")
	}

	if len(templateTypeAfterDeleteHooks) != 0 {
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
func (o *TemplateType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTemplateType(ctx, exec, o.TypeName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TemplateTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TemplateTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), templateTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"template_types\".* FROM \"vehicle_signal_decoding_api\".\"template_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, templateTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TemplateTypeSlice")
	}

	*o = slice

	return nil
}

// TemplateTypeExists checks if the TemplateType row exists.
func TemplateTypeExists(ctx context.Context, exec boil.ContextExecutor, typeName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"template_types\" where \"type_name\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, typeName)
	}
	row := exec.QueryRowContext(ctx, sql, typeName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if template_types exists")
	}

	return exists, nil
}

// Exists checks if the TemplateType row exists.
func (o *TemplateType) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TemplateTypeExists(ctx, exec, o.TypeName)
}
