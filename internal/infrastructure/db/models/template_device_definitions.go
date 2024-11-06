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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TemplateDeviceDefinition is an object representing the database table.
type TemplateDeviceDefinition struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeviceStyleID null.String `boil:"device_style_id" json:"device_style_id,omitempty" toml:"device_style_id" yaml:"device_style_id,omitempty"`
	TemplateName  string      `boil:"template_name" json:"template_name" toml:"template_name" yaml:"template_name"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DefinitionID  string      `boil:"definition_id" json:"definition_id" toml:"definition_id" yaml:"definition_id"`

	R *templateDeviceDefinitionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L templateDeviceDefinitionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TemplateDeviceDefinitionColumns = struct {
	ID            string
	DeviceStyleID string
	TemplateName  string
	CreatedAt     string
	UpdatedAt     string
	DefinitionID  string
}{
	ID:            "id",
	DeviceStyleID: "device_style_id",
	TemplateName:  "template_name",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DefinitionID:  "definition_id",
}

var TemplateDeviceDefinitionTableColumns = struct {
	ID            string
	DeviceStyleID string
	TemplateName  string
	CreatedAt     string
	UpdatedAt     string
	DefinitionID  string
}{
	ID:            "template_device_definitions.id",
	DeviceStyleID: "template_device_definitions.device_style_id",
	TemplateName:  "template_device_definitions.template_name",
	CreatedAt:     "template_device_definitions.created_at",
	UpdatedAt:     "template_device_definitions.updated_at",
	DefinitionID:  "template_device_definitions.definition_id",
}

// Generated where

var TemplateDeviceDefinitionWhere = struct {
	ID            whereHelperint64
	DeviceStyleID whereHelpernull_String
	TemplateName  whereHelperstring
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
	DefinitionID  whereHelperstring
}{
	ID:            whereHelperint64{field: "\"vehicle_signal_decoding_api\".\"template_device_definitions\".\"id\""},
	DeviceStyleID: whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"template_device_definitions\".\"device_style_id\""},
	TemplateName:  whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"template_device_definitions\".\"template_name\""},
	CreatedAt:     whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"template_device_definitions\".\"created_at\""},
	UpdatedAt:     whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"template_device_definitions\".\"updated_at\""},
	DefinitionID:  whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"template_device_definitions\".\"definition_id\""},
}

// TemplateDeviceDefinitionRels is where relationship names are stored.
var TemplateDeviceDefinitionRels = struct {
	TemplateNameTemplate string
}{
	TemplateNameTemplate: "TemplateNameTemplate",
}

// templateDeviceDefinitionR is where relationships are stored.
type templateDeviceDefinitionR struct {
	TemplateNameTemplate *Template `boil:"TemplateNameTemplate" json:"TemplateNameTemplate" toml:"TemplateNameTemplate" yaml:"TemplateNameTemplate"`
}

// NewStruct creates a new relationship struct
func (*templateDeviceDefinitionR) NewStruct() *templateDeviceDefinitionR {
	return &templateDeviceDefinitionR{}
}

func (r *templateDeviceDefinitionR) GetTemplateNameTemplate() *Template {
	if r == nil {
		return nil
	}
	return r.TemplateNameTemplate
}

// templateDeviceDefinitionL is where Load methods for each relationship are stored.
type templateDeviceDefinitionL struct{}

var (
	templateDeviceDefinitionAllColumns            = []string{"id", "device_style_id", "template_name", "created_at", "updated_at", "definition_id"}
	templateDeviceDefinitionColumnsWithoutDefault = []string{"template_name", "definition_id"}
	templateDeviceDefinitionColumnsWithDefault    = []string{"id", "device_style_id", "created_at", "updated_at"}
	templateDeviceDefinitionPrimaryKeyColumns     = []string{"id"}
	templateDeviceDefinitionGeneratedColumns      = []string{}
)

type (
	// TemplateDeviceDefinitionSlice is an alias for a slice of pointers to TemplateDeviceDefinition.
	// This should almost always be used instead of []TemplateDeviceDefinition.
	TemplateDeviceDefinitionSlice []*TemplateDeviceDefinition
	// TemplateDeviceDefinitionHook is the signature for custom TemplateDeviceDefinition hook methods
	TemplateDeviceDefinitionHook func(context.Context, boil.ContextExecutor, *TemplateDeviceDefinition) error

	templateDeviceDefinitionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	templateDeviceDefinitionType                 = reflect.TypeOf(&TemplateDeviceDefinition{})
	templateDeviceDefinitionMapping              = queries.MakeStructMapping(templateDeviceDefinitionType)
	templateDeviceDefinitionPrimaryKeyMapping, _ = queries.BindMapping(templateDeviceDefinitionType, templateDeviceDefinitionMapping, templateDeviceDefinitionPrimaryKeyColumns)
	templateDeviceDefinitionInsertCacheMut       sync.RWMutex
	templateDeviceDefinitionInsertCache          = make(map[string]insertCache)
	templateDeviceDefinitionUpdateCacheMut       sync.RWMutex
	templateDeviceDefinitionUpdateCache          = make(map[string]updateCache)
	templateDeviceDefinitionUpsertCacheMut       sync.RWMutex
	templateDeviceDefinitionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var templateDeviceDefinitionAfterSelectMu sync.Mutex
var templateDeviceDefinitionAfterSelectHooks []TemplateDeviceDefinitionHook

var templateDeviceDefinitionBeforeInsertMu sync.Mutex
var templateDeviceDefinitionBeforeInsertHooks []TemplateDeviceDefinitionHook
var templateDeviceDefinitionAfterInsertMu sync.Mutex
var templateDeviceDefinitionAfterInsertHooks []TemplateDeviceDefinitionHook

var templateDeviceDefinitionBeforeUpdateMu sync.Mutex
var templateDeviceDefinitionBeforeUpdateHooks []TemplateDeviceDefinitionHook
var templateDeviceDefinitionAfterUpdateMu sync.Mutex
var templateDeviceDefinitionAfterUpdateHooks []TemplateDeviceDefinitionHook

var templateDeviceDefinitionBeforeDeleteMu sync.Mutex
var templateDeviceDefinitionBeforeDeleteHooks []TemplateDeviceDefinitionHook
var templateDeviceDefinitionAfterDeleteMu sync.Mutex
var templateDeviceDefinitionAfterDeleteHooks []TemplateDeviceDefinitionHook

var templateDeviceDefinitionBeforeUpsertMu sync.Mutex
var templateDeviceDefinitionBeforeUpsertHooks []TemplateDeviceDefinitionHook
var templateDeviceDefinitionAfterUpsertMu sync.Mutex
var templateDeviceDefinitionAfterUpsertHooks []TemplateDeviceDefinitionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TemplateDeviceDefinition) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TemplateDeviceDefinition) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TemplateDeviceDefinition) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TemplateDeviceDefinition) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TemplateDeviceDefinition) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TemplateDeviceDefinition) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TemplateDeviceDefinition) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TemplateDeviceDefinition) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TemplateDeviceDefinition) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range templateDeviceDefinitionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTemplateDeviceDefinitionHook registers your hook function for all future operations.
func AddTemplateDeviceDefinitionHook(hookPoint boil.HookPoint, templateDeviceDefinitionHook TemplateDeviceDefinitionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		templateDeviceDefinitionAfterSelectMu.Lock()
		templateDeviceDefinitionAfterSelectHooks = append(templateDeviceDefinitionAfterSelectHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		templateDeviceDefinitionBeforeInsertMu.Lock()
		templateDeviceDefinitionBeforeInsertHooks = append(templateDeviceDefinitionBeforeInsertHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		templateDeviceDefinitionAfterInsertMu.Lock()
		templateDeviceDefinitionAfterInsertHooks = append(templateDeviceDefinitionAfterInsertHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		templateDeviceDefinitionBeforeUpdateMu.Lock()
		templateDeviceDefinitionBeforeUpdateHooks = append(templateDeviceDefinitionBeforeUpdateHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		templateDeviceDefinitionAfterUpdateMu.Lock()
		templateDeviceDefinitionAfterUpdateHooks = append(templateDeviceDefinitionAfterUpdateHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		templateDeviceDefinitionBeforeDeleteMu.Lock()
		templateDeviceDefinitionBeforeDeleteHooks = append(templateDeviceDefinitionBeforeDeleteHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		templateDeviceDefinitionAfterDeleteMu.Lock()
		templateDeviceDefinitionAfterDeleteHooks = append(templateDeviceDefinitionAfterDeleteHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		templateDeviceDefinitionBeforeUpsertMu.Lock()
		templateDeviceDefinitionBeforeUpsertHooks = append(templateDeviceDefinitionBeforeUpsertHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		templateDeviceDefinitionAfterUpsertMu.Lock()
		templateDeviceDefinitionAfterUpsertHooks = append(templateDeviceDefinitionAfterUpsertHooks, templateDeviceDefinitionHook)
		templateDeviceDefinitionAfterUpsertMu.Unlock()
	}
}

// One returns a single templateDeviceDefinition record from the query.
func (q templateDeviceDefinitionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TemplateDeviceDefinition, error) {
	o := &TemplateDeviceDefinition{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for template_device_definitions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TemplateDeviceDefinition records from the query.
func (q templateDeviceDefinitionQuery) All(ctx context.Context, exec boil.ContextExecutor) (TemplateDeviceDefinitionSlice, error) {
	var o []*TemplateDeviceDefinition

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TemplateDeviceDefinition slice")
	}

	if len(templateDeviceDefinitionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TemplateDeviceDefinition records in the query.
func (q templateDeviceDefinitionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count template_device_definitions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q templateDeviceDefinitionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if template_device_definitions exists")
	}

	return count > 0, nil
}

// TemplateNameTemplate pointed to by the foreign key.
func (o *TemplateDeviceDefinition) TemplateNameTemplate(mods ...qm.QueryMod) templateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"template_name\" = ?", o.TemplateName),
	}

	queryMods = append(queryMods, mods...)

	return Templates(queryMods...)
}

// LoadTemplateNameTemplate allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (templateDeviceDefinitionL) LoadTemplateNameTemplate(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTemplateDeviceDefinition interface{}, mods queries.Applicator) error {
	var slice []*TemplateDeviceDefinition
	var object *TemplateDeviceDefinition

	if singular {
		var ok bool
		object, ok = maybeTemplateDeviceDefinition.(*TemplateDeviceDefinition)
		if !ok {
			object = new(TemplateDeviceDefinition)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTemplateDeviceDefinition)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTemplateDeviceDefinition))
			}
		}
	} else {
		s, ok := maybeTemplateDeviceDefinition.(*[]*TemplateDeviceDefinition)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTemplateDeviceDefinition)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTemplateDeviceDefinition))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &templateDeviceDefinitionR{}
		}
		args[object.TemplateName] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &templateDeviceDefinitionR{}
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
		foreign.R.TemplateNameTemplateDeviceDefinitions = append(foreign.R.TemplateNameTemplateDeviceDefinitions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TemplateName == foreign.TemplateName {
				local.R.TemplateNameTemplate = foreign
				if foreign.R == nil {
					foreign.R = &templateR{}
				}
				foreign.R.TemplateNameTemplateDeviceDefinitions = append(foreign.R.TemplateNameTemplateDeviceDefinitions, local)
				break
			}
		}
	}

	return nil
}

// SetTemplateNameTemplate of the templateDeviceDefinition to the related item.
// Sets o.R.TemplateNameTemplate to related.
// Adds o to related.R.TemplateNameTemplateDeviceDefinitions.
func (o *TemplateDeviceDefinition) SetTemplateNameTemplate(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Template) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"vehicle_signal_decoding_api\".\"template_device_definitions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"template_name"}),
		strmangle.WhereClause("\"", "\"", 2, templateDeviceDefinitionPrimaryKeyColumns),
	)
	values := []interface{}{related.TemplateName, o.ID}

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
		o.R = &templateDeviceDefinitionR{
			TemplateNameTemplate: related,
		}
	} else {
		o.R.TemplateNameTemplate = related
	}

	if related.R == nil {
		related.R = &templateR{
			TemplateNameTemplateDeviceDefinitions: TemplateDeviceDefinitionSlice{o},
		}
	} else {
		related.R.TemplateNameTemplateDeviceDefinitions = append(related.R.TemplateNameTemplateDeviceDefinitions, o)
	}

	return nil
}

// TemplateDeviceDefinitions retrieves all the records using an executor.
func TemplateDeviceDefinitions(mods ...qm.QueryMod) templateDeviceDefinitionQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"template_device_definitions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"template_device_definitions\".*"})
	}

	return templateDeviceDefinitionQuery{q}
}

// FindTemplateDeviceDefinition retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTemplateDeviceDefinition(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TemplateDeviceDefinition, error) {
	templateDeviceDefinitionObj := &TemplateDeviceDefinition{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"template_device_definitions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, templateDeviceDefinitionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from template_device_definitions")
	}

	if err = templateDeviceDefinitionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return templateDeviceDefinitionObj, err
	}

	return templateDeviceDefinitionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TemplateDeviceDefinition) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no template_device_definitions provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(templateDeviceDefinitionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	templateDeviceDefinitionInsertCacheMut.RLock()
	cache, cached := templateDeviceDefinitionInsertCache[key]
	templateDeviceDefinitionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			templateDeviceDefinitionAllColumns,
			templateDeviceDefinitionColumnsWithDefault,
			templateDeviceDefinitionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(templateDeviceDefinitionType, templateDeviceDefinitionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(templateDeviceDefinitionType, templateDeviceDefinitionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"template_device_definitions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"template_device_definitions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into template_device_definitions")
	}

	if !cached {
		templateDeviceDefinitionInsertCacheMut.Lock()
		templateDeviceDefinitionInsertCache[key] = cache
		templateDeviceDefinitionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TemplateDeviceDefinition.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TemplateDeviceDefinition) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	templateDeviceDefinitionUpdateCacheMut.RLock()
	cache, cached := templateDeviceDefinitionUpdateCache[key]
	templateDeviceDefinitionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			templateDeviceDefinitionAllColumns,
			templateDeviceDefinitionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update template_device_definitions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"template_device_definitions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, templateDeviceDefinitionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(templateDeviceDefinitionType, templateDeviceDefinitionMapping, append(wl, templateDeviceDefinitionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update template_device_definitions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for template_device_definitions")
	}

	if !cached {
		templateDeviceDefinitionUpdateCacheMut.Lock()
		templateDeviceDefinitionUpdateCache[key] = cache
		templateDeviceDefinitionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q templateDeviceDefinitionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for template_device_definitions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for template_device_definitions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TemplateDeviceDefinitionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), templateDeviceDefinitionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"template_device_definitions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, templateDeviceDefinitionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in templateDeviceDefinition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all templateDeviceDefinition")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TemplateDeviceDefinition) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no template_device_definitions provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(templateDeviceDefinitionColumnsWithDefault, o)

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

	templateDeviceDefinitionUpsertCacheMut.RLock()
	cache, cached := templateDeviceDefinitionUpsertCache[key]
	templateDeviceDefinitionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			templateDeviceDefinitionAllColumns,
			templateDeviceDefinitionColumnsWithDefault,
			templateDeviceDefinitionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			templateDeviceDefinitionAllColumns,
			templateDeviceDefinitionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert template_device_definitions, could not build update column list")
		}

		ret := strmangle.SetComplement(templateDeviceDefinitionAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(templateDeviceDefinitionPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert template_device_definitions, could not build conflict column list")
			}

			conflict = make([]string, len(templateDeviceDefinitionPrimaryKeyColumns))
			copy(conflict, templateDeviceDefinitionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"template_device_definitions\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(templateDeviceDefinitionType, templateDeviceDefinitionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(templateDeviceDefinitionType, templateDeviceDefinitionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert template_device_definitions")
	}

	if !cached {
		templateDeviceDefinitionUpsertCacheMut.Lock()
		templateDeviceDefinitionUpsertCache[key] = cache
		templateDeviceDefinitionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TemplateDeviceDefinition record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TemplateDeviceDefinition) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TemplateDeviceDefinition provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), templateDeviceDefinitionPrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"template_device_definitions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from template_device_definitions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for template_device_definitions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q templateDeviceDefinitionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no templateDeviceDefinitionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from template_device_definitions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for template_device_definitions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TemplateDeviceDefinitionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(templateDeviceDefinitionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), templateDeviceDefinitionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"template_device_definitions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, templateDeviceDefinitionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from templateDeviceDefinition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for template_device_definitions")
	}

	if len(templateDeviceDefinitionAfterDeleteHooks) != 0 {
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
func (o *TemplateDeviceDefinition) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTemplateDeviceDefinition(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TemplateDeviceDefinitionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TemplateDeviceDefinitionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), templateDeviceDefinitionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"template_device_definitions\".* FROM \"vehicle_signal_decoding_api\".\"template_device_definitions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, templateDeviceDefinitionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TemplateDeviceDefinitionSlice")
	}

	*o = slice

	return nil
}

// TemplateDeviceDefinitionExists checks if the TemplateDeviceDefinition row exists.
func TemplateDeviceDefinitionExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"template_device_definitions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if template_device_definitions exists")
	}

	return exists, nil
}

// Exists checks if the TemplateDeviceDefinition row exists.
func (o *TemplateDeviceDefinition) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TemplateDeviceDefinitionExists(ctx, exec, o.ID)
}
