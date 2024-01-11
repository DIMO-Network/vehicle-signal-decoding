// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// DeviceSetting is an object representing the database table.
type DeviceSetting struct {
	TemplateName null.String `boil:"template_name" json:"template_name,omitempty" toml:"template_name" yaml:"template_name,omitempty"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Settings     null.JSON   `boil:"settings" json:"settings,omitempty" toml:"settings" yaml:"settings,omitempty"`
	Name         string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Powertrain   string      `boil:"powertrain" json:"powertrain" toml:"powertrain" yaml:"powertrain"`

	R *deviceSettingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deviceSettingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeviceSettingColumns = struct {
	TemplateName string
	CreatedAt    string
	UpdatedAt    string
	Settings     string
	Name         string
	Powertrain   string
}{
	TemplateName: "template_name",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Settings:     "settings",
	Name:         "name",
	Powertrain:   "powertrain",
}

var DeviceSettingTableColumns = struct {
	TemplateName string
	CreatedAt    string
	UpdatedAt    string
	Settings     string
	Name         string
	Powertrain   string
}{
	TemplateName: "device_settings.template_name",
	CreatedAt:    "device_settings.created_at",
	UpdatedAt:    "device_settings.updated_at",
	Settings:     "device_settings.settings",
	Name:         "device_settings.name",
	Powertrain:   "device_settings.powertrain",
}

// Generated where

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var DeviceSettingWhere = struct {
	TemplateName whereHelpernull_String
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
	Settings     whereHelpernull_JSON
	Name         whereHelperstring
	Powertrain   whereHelperstring
}{
	TemplateName: whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"device_settings\".\"template_name\""},
	CreatedAt:    whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"device_settings\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"device_settings\".\"updated_at\""},
	Settings:     whereHelpernull_JSON{field: "\"vehicle_signal_decoding_api\".\"device_settings\".\"settings\""},
	Name:         whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"device_settings\".\"name\""},
	Powertrain:   whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"device_settings\".\"powertrain\""},
}

// DeviceSettingRels is where relationship names are stored.
var DeviceSettingRels = struct {
	TemplateNameTemplate string
}{
	TemplateNameTemplate: "TemplateNameTemplate",
}

// deviceSettingR is where relationships are stored.
type deviceSettingR struct {
	TemplateNameTemplate *Template `boil:"TemplateNameTemplate" json:"TemplateNameTemplate" toml:"TemplateNameTemplate" yaml:"TemplateNameTemplate"`
}

// NewStruct creates a new relationship struct
func (*deviceSettingR) NewStruct() *deviceSettingR {
	return &deviceSettingR{}
}

func (r *deviceSettingR) GetTemplateNameTemplate() *Template {
	if r == nil {
		return nil
	}
	return r.TemplateNameTemplate
}

// deviceSettingL is where Load methods for each relationship are stored.
type deviceSettingL struct{}

var (
	deviceSettingAllColumns            = []string{"template_name", "created_at", "updated_at", "settings", "name", "powertrain"}
	deviceSettingColumnsWithoutDefault = []string{"name"}
	deviceSettingColumnsWithDefault    = []string{"template_name", "created_at", "updated_at", "settings", "powertrain"}
	deviceSettingPrimaryKeyColumns     = []string{"name"}
	deviceSettingGeneratedColumns      = []string{}
)

type (
	// DeviceSettingSlice is an alias for a slice of pointers to DeviceSetting.
	// This should almost always be used instead of []DeviceSetting.
	DeviceSettingSlice []*DeviceSetting
	// DeviceSettingHook is the signature for custom DeviceSetting hook methods
	DeviceSettingHook func(context.Context, boil.ContextExecutor, *DeviceSetting) error

	deviceSettingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deviceSettingType                 = reflect.TypeOf(&DeviceSetting{})
	deviceSettingMapping              = queries.MakeStructMapping(deviceSettingType)
	deviceSettingPrimaryKeyMapping, _ = queries.BindMapping(deviceSettingType, deviceSettingMapping, deviceSettingPrimaryKeyColumns)
	deviceSettingInsertCacheMut       sync.RWMutex
	deviceSettingInsertCache          = make(map[string]insertCache)
	deviceSettingUpdateCacheMut       sync.RWMutex
	deviceSettingUpdateCache          = make(map[string]updateCache)
	deviceSettingUpsertCacheMut       sync.RWMutex
	deviceSettingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var deviceSettingAfterSelectHooks []DeviceSettingHook

var deviceSettingBeforeInsertHooks []DeviceSettingHook
var deviceSettingAfterInsertHooks []DeviceSettingHook

var deviceSettingBeforeUpdateHooks []DeviceSettingHook
var deviceSettingAfterUpdateHooks []DeviceSettingHook

var deviceSettingBeforeDeleteHooks []DeviceSettingHook
var deviceSettingAfterDeleteHooks []DeviceSettingHook

var deviceSettingBeforeUpsertHooks []DeviceSettingHook
var deviceSettingAfterUpsertHooks []DeviceSettingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DeviceSetting) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DeviceSetting) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DeviceSetting) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DeviceSetting) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DeviceSetting) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DeviceSetting) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DeviceSetting) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DeviceSetting) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DeviceSetting) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceSettingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDeviceSettingHook registers your hook function for all future operations.
func AddDeviceSettingHook(hookPoint boil.HookPoint, deviceSettingHook DeviceSettingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		deviceSettingAfterSelectHooks = append(deviceSettingAfterSelectHooks, deviceSettingHook)
	case boil.BeforeInsertHook:
		deviceSettingBeforeInsertHooks = append(deviceSettingBeforeInsertHooks, deviceSettingHook)
	case boil.AfterInsertHook:
		deviceSettingAfterInsertHooks = append(deviceSettingAfterInsertHooks, deviceSettingHook)
	case boil.BeforeUpdateHook:
		deviceSettingBeforeUpdateHooks = append(deviceSettingBeforeUpdateHooks, deviceSettingHook)
	case boil.AfterUpdateHook:
		deviceSettingAfterUpdateHooks = append(deviceSettingAfterUpdateHooks, deviceSettingHook)
	case boil.BeforeDeleteHook:
		deviceSettingBeforeDeleteHooks = append(deviceSettingBeforeDeleteHooks, deviceSettingHook)
	case boil.AfterDeleteHook:
		deviceSettingAfterDeleteHooks = append(deviceSettingAfterDeleteHooks, deviceSettingHook)
	case boil.BeforeUpsertHook:
		deviceSettingBeforeUpsertHooks = append(deviceSettingBeforeUpsertHooks, deviceSettingHook)
	case boil.AfterUpsertHook:
		deviceSettingAfterUpsertHooks = append(deviceSettingAfterUpsertHooks, deviceSettingHook)
	}
}

// One returns a single deviceSetting record from the query.
func (q deviceSettingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DeviceSetting, error) {
	o := &DeviceSetting{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for device_settings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DeviceSetting records from the query.
func (q deviceSettingQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeviceSettingSlice, error) {
	var o []*DeviceSetting

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DeviceSetting slice")
	}

	if len(deviceSettingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DeviceSetting records in the query.
func (q deviceSettingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count device_settings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deviceSettingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if device_settings exists")
	}

	return count > 0, nil
}

// TemplateNameTemplate pointed to by the foreign key.
func (o *DeviceSetting) TemplateNameTemplate(mods ...qm.QueryMod) templateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"template_name\" = ?", o.TemplateName),
	}

	queryMods = append(queryMods, mods...)

	return Templates(queryMods...)
}

// LoadTemplateNameTemplate allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (deviceSettingL) LoadTemplateNameTemplate(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDeviceSetting interface{}, mods queries.Applicator) error {
	var slice []*DeviceSetting
	var object *DeviceSetting

	if singular {
		var ok bool
		object, ok = maybeDeviceSetting.(*DeviceSetting)
		if !ok {
			object = new(DeviceSetting)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDeviceSetting)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDeviceSetting))
			}
		}
	} else {
		s, ok := maybeDeviceSetting.(*[]*DeviceSetting)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDeviceSetting)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDeviceSetting))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &deviceSettingR{}
		}
		if !queries.IsNil(object.TemplateName) {
			args = append(args, object.TemplateName)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &deviceSettingR{}
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

	if len(deviceSettingAfterSelectHooks) != 0 {
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
		foreign.R.TemplateNameDeviceSettings = append(foreign.R.TemplateNameDeviceSettings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TemplateName, foreign.TemplateName) {
				local.R.TemplateNameTemplate = foreign
				if foreign.R == nil {
					foreign.R = &templateR{}
				}
				foreign.R.TemplateNameDeviceSettings = append(foreign.R.TemplateNameDeviceSettings, local)
				break
			}
		}
	}

	return nil
}

// SetTemplateNameTemplate of the deviceSetting to the related item.
// Sets o.R.TemplateNameTemplate to related.
// Adds o to related.R.TemplateNameDeviceSettings.
func (o *DeviceSetting) SetTemplateNameTemplate(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Template) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"vehicle_signal_decoding_api\".\"device_settings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"template_name"}),
		strmangle.WhereClause("\"", "\"", 2, deviceSettingPrimaryKeyColumns),
	)
	values := []interface{}{related.TemplateName, o.Name}

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
		o.R = &deviceSettingR{
			TemplateNameTemplate: related,
		}
	} else {
		o.R.TemplateNameTemplate = related
	}

	if related.R == nil {
		related.R = &templateR{
			TemplateNameDeviceSettings: DeviceSettingSlice{o},
		}
	} else {
		related.R.TemplateNameDeviceSettings = append(related.R.TemplateNameDeviceSettings, o)
	}

	return nil
}

// RemoveTemplateNameTemplate relationship.
// Sets o.R.TemplateNameTemplate to nil.
// Removes o from all passed in related items' relationships struct.
func (o *DeviceSetting) RemoveTemplateNameTemplate(ctx context.Context, exec boil.ContextExecutor, related *Template) error {
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

	for i, ri := range related.R.TemplateNameDeviceSettings {
		if queries.Equal(o.TemplateName, ri.TemplateName) {
			continue
		}

		ln := len(related.R.TemplateNameDeviceSettings)
		if ln > 1 && i < ln-1 {
			related.R.TemplateNameDeviceSettings[i] = related.R.TemplateNameDeviceSettings[ln-1]
		}
		related.R.TemplateNameDeviceSettings = related.R.TemplateNameDeviceSettings[:ln-1]
		break
	}
	return nil
}

// DeviceSettings retrieves all the records using an executor.
func DeviceSettings(mods ...qm.QueryMod) deviceSettingQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"device_settings\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"device_settings\".*"})
	}

	return deviceSettingQuery{q}
}

// FindDeviceSetting retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDeviceSetting(ctx context.Context, exec boil.ContextExecutor, name string, selectCols ...string) (*DeviceSetting, error) {
	deviceSettingObj := &DeviceSetting{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"device_settings\" where \"name\"=$1", sel,
	)

	q := queries.Raw(query, name)

	err := q.Bind(ctx, exec, deviceSettingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from device_settings")
	}

	if err = deviceSettingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return deviceSettingObj, err
	}

	return deviceSettingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DeviceSetting) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no device_settings provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(deviceSettingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deviceSettingInsertCacheMut.RLock()
	cache, cached := deviceSettingInsertCache[key]
	deviceSettingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deviceSettingAllColumns,
			deviceSettingColumnsWithDefault,
			deviceSettingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deviceSettingType, deviceSettingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deviceSettingType, deviceSettingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"device_settings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"device_settings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into device_settings")
	}

	if !cached {
		deviceSettingInsertCacheMut.Lock()
		deviceSettingInsertCache[key] = cache
		deviceSettingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DeviceSetting.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DeviceSetting) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	deviceSettingUpdateCacheMut.RLock()
	cache, cached := deviceSettingUpdateCache[key]
	deviceSettingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deviceSettingAllColumns,
			deviceSettingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update device_settings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"device_settings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, deviceSettingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deviceSettingType, deviceSettingMapping, append(wl, deviceSettingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update device_settings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for device_settings")
	}

	if !cached {
		deviceSettingUpdateCacheMut.Lock()
		deviceSettingUpdateCache[key] = cache
		deviceSettingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q deviceSettingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for device_settings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for device_settings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DeviceSettingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"device_settings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, deviceSettingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in deviceSetting slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all deviceSetting")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DeviceSetting) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no device_settings provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(deviceSettingColumnsWithDefault, o)

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

	deviceSettingUpsertCacheMut.RLock()
	cache, cached := deviceSettingUpsertCache[key]
	deviceSettingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			deviceSettingAllColumns,
			deviceSettingColumnsWithDefault,
			deviceSettingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			deviceSettingAllColumns,
			deviceSettingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert device_settings, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(deviceSettingPrimaryKeyColumns))
			copy(conflict, deviceSettingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"device_settings\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(deviceSettingType, deviceSettingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deviceSettingType, deviceSettingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert device_settings")
	}

	if !cached {
		deviceSettingUpsertCacheMut.Lock()
		deviceSettingUpsertCache[key] = cache
		deviceSettingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DeviceSetting record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DeviceSetting) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DeviceSetting provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deviceSettingPrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"device_settings\" WHERE \"name\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from device_settings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for device_settings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deviceSettingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no deviceSettingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from device_settings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for device_settings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DeviceSettingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(deviceSettingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"device_settings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, deviceSettingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deviceSetting slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for device_settings")
	}

	if len(deviceSettingAfterDeleteHooks) != 0 {
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
func (o *DeviceSetting) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDeviceSetting(ctx, exec, o.Name)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DeviceSettingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeviceSettingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"device_settings\".* FROM \"vehicle_signal_decoding_api\".\"device_settings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, deviceSettingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DeviceSettingSlice")
	}

	*o = slice

	return nil
}

// DeviceSettingExists checks if the DeviceSetting row exists.
func DeviceSettingExists(ctx context.Context, exec boil.ContextExecutor, name string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"device_settings\" where \"name\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, name)
	}
	row := exec.QueryRowContext(ctx, sql, name)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if device_settings exists")
	}

	return exists, nil
}
