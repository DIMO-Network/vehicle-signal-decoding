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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PidConfig is an object representing the database table.
type PidConfig struct {
	ID              int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TemplateName    null.String `boil:"template_name" json:"template_name,omitempty" toml:"template_name" yaml:"template_name,omitempty"`
	Header          []byte      `boil:"header" json:"header" toml:"header" yaml:"header"`
	Mode            []byte      `boil:"mode" json:"mode" toml:"mode" yaml:"mode"`
	Pid             []byte      `boil:"pid" json:"pid" toml:"pid" yaml:"pid"`
	Formula         string      `boil:"formula" json:"formula" toml:"formula" yaml:"formula"`
	IntervalSeconds int         `boil:"interval_seconds" json:"interval_seconds" toml:"interval_seconds" yaml:"interval_seconds"`
	Version         null.String `boil:"version" json:"version,omitempty" toml:"version" yaml:"version,omitempty"`
	CreatedAt       time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *pidConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pidConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PidConfigColumns = struct {
	ID              string
	TemplateName    string
	Header          string
	Mode            string
	Pid             string
	Formula         string
	IntervalSeconds string
	Version         string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	TemplateName:    "template_name",
	Header:          "header",
	Mode:            "mode",
	Pid:             "pid",
	Formula:         "formula",
	IntervalSeconds: "interval_seconds",
	Version:         "version",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

var PidConfigTableColumns = struct {
	ID              string
	TemplateName    string
	Header          string
	Mode            string
	Pid             string
	Formula         string
	IntervalSeconds string
	Version         string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "pid_configs.id",
	TemplateName:    "pid_configs.template_name",
	Header:          "pid_configs.header",
	Mode:            "pid_configs.mode",
	Pid:             "pid_configs.pid",
	Formula:         "pid_configs.formula",
	IntervalSeconds: "pid_configs.interval_seconds",
	Version:         "pid_configs.version",
	CreatedAt:       "pid_configs.created_at",
	UpdatedAt:       "pid_configs.updated_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var PidConfigWhere = struct {
	ID              whereHelperint64
	TemplateName    whereHelpernull_String
	Header          whereHelper__byte
	Mode            whereHelper__byte
	Pid             whereHelper__byte
	Formula         whereHelperstring
	IntervalSeconds whereHelperint
	Version         whereHelpernull_String
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
}{
	ID:              whereHelperint64{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"id\""},
	TemplateName:    whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"template_name\""},
	Header:          whereHelper__byte{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"header\""},
	Mode:            whereHelper__byte{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"mode\""},
	Pid:             whereHelper__byte{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"pid\""},
	Formula:         whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"formula\""},
	IntervalSeconds: whereHelperint{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"interval_seconds\""},
	Version:         whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"version\""},
	CreatedAt:       whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"created_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"pid_configs\".\"updated_at\""},
}

// PidConfigRels is where relationship names are stored.
var PidConfigRels = struct {
	TemplateNameTemplate string
}{
	TemplateNameTemplate: "TemplateNameTemplate",
}

// pidConfigR is where relationships are stored.
type pidConfigR struct {
	TemplateNameTemplate *Template `boil:"TemplateNameTemplate" json:"TemplateNameTemplate" toml:"TemplateNameTemplate" yaml:"TemplateNameTemplate"`
}

// NewStruct creates a new relationship struct
func (*pidConfigR) NewStruct() *pidConfigR {
	return &pidConfigR{}
}

func (r *pidConfigR) GetTemplateNameTemplate() *Template {
	if r == nil {
		return nil
	}
	return r.TemplateNameTemplate
}

// pidConfigL is where Load methods for each relationship are stored.
type pidConfigL struct{}

var (
	pidConfigAllColumns            = []string{"id", "template_name", "header", "mode", "pid", "formula", "interval_seconds", "version", "created_at", "updated_at"}
	pidConfigColumnsWithoutDefault = []string{"header", "mode", "pid", "formula", "interval_seconds"}
	pidConfigColumnsWithDefault    = []string{"id", "template_name", "version", "created_at", "updated_at"}
	pidConfigPrimaryKeyColumns     = []string{"id"}
	pidConfigGeneratedColumns      = []string{}
)

type (
	// PidConfigSlice is an alias for a slice of pointers to PidConfig.
	// This should almost always be used instead of []PidConfig.
	PidConfigSlice []*PidConfig
	// PidConfigHook is the signature for custom PidConfig hook methods
	PidConfigHook func(context.Context, boil.ContextExecutor, *PidConfig) error

	pidConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pidConfigType                 = reflect.TypeOf(&PidConfig{})
	pidConfigMapping              = queries.MakeStructMapping(pidConfigType)
	pidConfigPrimaryKeyMapping, _ = queries.BindMapping(pidConfigType, pidConfigMapping, pidConfigPrimaryKeyColumns)
	pidConfigInsertCacheMut       sync.RWMutex
	pidConfigInsertCache          = make(map[string]insertCache)
	pidConfigUpdateCacheMut       sync.RWMutex
	pidConfigUpdateCache          = make(map[string]updateCache)
	pidConfigUpsertCacheMut       sync.RWMutex
	pidConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var pidConfigAfterSelectHooks []PidConfigHook

var pidConfigBeforeInsertHooks []PidConfigHook
var pidConfigAfterInsertHooks []PidConfigHook

var pidConfigBeforeUpdateHooks []PidConfigHook
var pidConfigAfterUpdateHooks []PidConfigHook

var pidConfigBeforeDeleteHooks []PidConfigHook
var pidConfigAfterDeleteHooks []PidConfigHook

var pidConfigBeforeUpsertHooks []PidConfigHook
var pidConfigAfterUpsertHooks []PidConfigHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PidConfig) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PidConfig) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PidConfig) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PidConfig) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PidConfig) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PidConfig) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PidConfig) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PidConfig) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PidConfig) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pidConfigAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPidConfigHook registers your hook function for all future operations.
func AddPidConfigHook(hookPoint boil.HookPoint, pidConfigHook PidConfigHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		pidConfigAfterSelectHooks = append(pidConfigAfterSelectHooks, pidConfigHook)
	case boil.BeforeInsertHook:
		pidConfigBeforeInsertHooks = append(pidConfigBeforeInsertHooks, pidConfigHook)
	case boil.AfterInsertHook:
		pidConfigAfterInsertHooks = append(pidConfigAfterInsertHooks, pidConfigHook)
	case boil.BeforeUpdateHook:
		pidConfigBeforeUpdateHooks = append(pidConfigBeforeUpdateHooks, pidConfigHook)
	case boil.AfterUpdateHook:
		pidConfigAfterUpdateHooks = append(pidConfigAfterUpdateHooks, pidConfigHook)
	case boil.BeforeDeleteHook:
		pidConfigBeforeDeleteHooks = append(pidConfigBeforeDeleteHooks, pidConfigHook)
	case boil.AfterDeleteHook:
		pidConfigAfterDeleteHooks = append(pidConfigAfterDeleteHooks, pidConfigHook)
	case boil.BeforeUpsertHook:
		pidConfigBeforeUpsertHooks = append(pidConfigBeforeUpsertHooks, pidConfigHook)
	case boil.AfterUpsertHook:
		pidConfigAfterUpsertHooks = append(pidConfigAfterUpsertHooks, pidConfigHook)
	}
}

// One returns a single pidConfig record from the query.
func (q pidConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PidConfig, error) {
	o := &PidConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pid_configs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PidConfig records from the query.
func (q pidConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (PidConfigSlice, error) {
	var o []*PidConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PidConfig slice")
	}

	if len(pidConfigAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PidConfig records in the query.
func (q pidConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pid_configs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pidConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pid_configs exists")
	}

	return count > 0, nil
}

// TemplateNameTemplate pointed to by the foreign key.
func (o *PidConfig) TemplateNameTemplate(mods ...qm.QueryMod) templateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"template_name\" = ?", o.TemplateName),
	}

	queryMods = append(queryMods, mods...)

	return Templates(queryMods...)
}

// LoadTemplateNameTemplate allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pidConfigL) LoadTemplateNameTemplate(ctx context.Context, e boil.ContextExecutor, singular bool, maybePidConfig interface{}, mods queries.Applicator) error {
	var slice []*PidConfig
	var object *PidConfig

	if singular {
		var ok bool
		object, ok = maybePidConfig.(*PidConfig)
		if !ok {
			object = new(PidConfig)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePidConfig)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePidConfig))
			}
		}
	} else {
		s, ok := maybePidConfig.(*[]*PidConfig)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePidConfig)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePidConfig))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pidConfigR{}
		}
		if !queries.IsNil(object.TemplateName) {
			args = append(args, object.TemplateName)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pidConfigR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.TemplateName) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.TemplateName) {
				args = append(args, obj.TemplateName)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`vehicle_signal_decoding_api.templates`),
		qm.WhereIn(`vehicle_signal_decoding_api.templates.template_name in ?`, args...),
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
		foreign.R.TemplateNamePidConfigs = append(foreign.R.TemplateNamePidConfigs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TemplateName, foreign.TemplateName) {
				local.R.TemplateNameTemplate = foreign
				if foreign.R == nil {
					foreign.R = &templateR{}
				}
				foreign.R.TemplateNamePidConfigs = append(foreign.R.TemplateNamePidConfigs, local)
				break
			}
		}
	}

	return nil
}

// SetTemplateNameTemplate of the pidConfig to the related item.
// Sets o.R.TemplateNameTemplate to related.
// Adds o to related.R.TemplateNamePidConfigs.
func (o *PidConfig) SetTemplateNameTemplate(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Template) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"vehicle_signal_decoding_api\".\"pid_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"template_name"}),
		strmangle.WhereClause("\"", "\"", 2, pidConfigPrimaryKeyColumns),
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

	queries.Assign(&o.TemplateName, related.TemplateName)
	if o.R == nil {
		o.R = &pidConfigR{
			TemplateNameTemplate: related,
		}
	} else {
		o.R.TemplateNameTemplate = related
	}

	if related.R == nil {
		related.R = &templateR{
			TemplateNamePidConfigs: PidConfigSlice{o},
		}
	} else {
		related.R.TemplateNamePidConfigs = append(related.R.TemplateNamePidConfigs, o)
	}

	return nil
}

// RemoveTemplateNameTemplate relationship.
// Sets o.R.TemplateNameTemplate to nil.
// Removes o from all passed in related items' relationships struct.
func (o *PidConfig) RemoveTemplateNameTemplate(ctx context.Context, exec boil.ContextExecutor, related *Template) error {
	var err error

	queries.SetScanner(&o.TemplateName, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("template_name")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.TemplateNameTemplate = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TemplateNamePidConfigs {
		if queries.Equal(o.TemplateName, ri.TemplateName) {
			continue
		}

		ln := len(related.R.TemplateNamePidConfigs)
		if ln > 1 && i < ln-1 {
			related.R.TemplateNamePidConfigs[i] = related.R.TemplateNamePidConfigs[ln-1]
		}
		related.R.TemplateNamePidConfigs = related.R.TemplateNamePidConfigs[:ln-1]
		break
	}
	return nil
}

// PidConfigs retrieves all the records using an executor.
func PidConfigs(mods ...qm.QueryMod) pidConfigQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"pid_configs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"pid_configs\".*"})
	}

	return pidConfigQuery{q}
}

// FindPidConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPidConfig(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PidConfig, error) {
	pidConfigObj := &PidConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"pid_configs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pidConfigObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pid_configs")
	}

	if err = pidConfigObj.doAfterSelectHooks(ctx, exec); err != nil {
		return pidConfigObj, err
	}

	return pidConfigObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PidConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pid_configs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(pidConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pidConfigInsertCacheMut.RLock()
	cache, cached := pidConfigInsertCache[key]
	pidConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pidConfigAllColumns,
			pidConfigColumnsWithDefault,
			pidConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pidConfigType, pidConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pidConfigType, pidConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"pid_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"pid_configs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into pid_configs")
	}

	if !cached {
		pidConfigInsertCacheMut.Lock()
		pidConfigInsertCache[key] = cache
		pidConfigInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PidConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PidConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pidConfigUpdateCacheMut.RLock()
	cache, cached := pidConfigUpdateCache[key]
	pidConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pidConfigAllColumns,
			pidConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pid_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"pid_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pidConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pidConfigType, pidConfigMapping, append(wl, pidConfigPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update pid_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pid_configs")
	}

	if !cached {
		pidConfigUpdateCacheMut.Lock()
		pidConfigUpdateCache[key] = cache
		pidConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pidConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pid_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pid_configs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PidConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pidConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"pid_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pidConfigPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pidConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pidConfig")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PidConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pid_configs provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(pidConfigColumnsWithDefault, o)

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

	pidConfigUpsertCacheMut.RLock()
	cache, cached := pidConfigUpsertCache[key]
	pidConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pidConfigAllColumns,
			pidConfigColumnsWithDefault,
			pidConfigColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			pidConfigAllColumns,
			pidConfigPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pid_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pidConfigPrimaryKeyColumns))
			copy(conflict, pidConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"pid_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pidConfigType, pidConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pidConfigType, pidConfigMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert pid_configs")
	}

	if !cached {
		pidConfigUpsertCacheMut.Lock()
		pidConfigUpsertCache[key] = cache
		pidConfigUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PidConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PidConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PidConfig provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pidConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"pid_configs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pid_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pid_configs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pidConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pidConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pid_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pid_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PidConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(pidConfigBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pidConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"pid_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pidConfigPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pidConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pid_configs")
	}

	if len(pidConfigAfterDeleteHooks) != 0 {
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
func (o *PidConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPidConfig(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PidConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PidConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pidConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"pid_configs\".* FROM \"vehicle_signal_decoding_api\".\"pid_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pidConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PidConfigSlice")
	}

	*o = slice

	return nil
}

// PidConfigExists checks if the PidConfig row exists.
func PidConfigExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"pid_configs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pid_configs exists")
	}

	return exists, nil
}

// Exists checks if the PidConfig row exists.
func (o *PidConfig) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PidConfigExists(ctx, exec, o.ID)
}
