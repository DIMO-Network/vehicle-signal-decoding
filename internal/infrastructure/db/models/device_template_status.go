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

// DeviceTemplateStatus is an object representing the database table.
type DeviceTemplateStatus struct {
	DeviceEthAddr       []byte      `boil:"device_eth_addr" json:"device_eth_addr" toml:"device_eth_addr" yaml:"device_eth_addr"`
	TemplateDBCURL      null.String `boil:"template_dbc_url" json:"template_dbc_url,omitempty" toml:"template_dbc_url" yaml:"template_dbc_url,omitempty"`
	TemplatePidURL      null.String `boil:"template_pid_url" json:"template_pid_url,omitempty" toml:"template_pid_url" yaml:"template_pid_url,omitempty"`
	TemplateSettingsURL null.String `boil:"template_settings_url" json:"template_settings_url,omitempty" toml:"template_settings_url" yaml:"template_settings_url,omitempty"`
	FirmwareVersion     null.String `boil:"firmware_version" json:"firmware_version,omitempty" toml:"firmware_version" yaml:"firmware_version,omitempty"`
	CreatedAt           time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt           time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *deviceTemplateStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deviceTemplateStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeviceTemplateStatusColumns = struct {
	DeviceEthAddr       string
	TemplateDBCURL      string
	TemplatePidURL      string
	TemplateSettingsURL string
	FirmwareVersion     string
	CreatedAt           string
	UpdatedAt           string
}{
	DeviceEthAddr:       "device_eth_addr",
	TemplateDBCURL:      "template_dbc_url",
	TemplatePidURL:      "template_pid_url",
	TemplateSettingsURL: "template_settings_url",
	FirmwareVersion:     "firmware_version",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

var DeviceTemplateStatusTableColumns = struct {
	DeviceEthAddr       string
	TemplateDBCURL      string
	TemplatePidURL      string
	TemplateSettingsURL string
	FirmwareVersion     string
	CreatedAt           string
	UpdatedAt           string
}{
	DeviceEthAddr:       "device_template_status.device_eth_addr",
	TemplateDBCURL:      "device_template_status.template_dbc_url",
	TemplatePidURL:      "device_template_status.template_pid_url",
	TemplateSettingsURL: "device_template_status.template_settings_url",
	FirmwareVersion:     "device_template_status.firmware_version",
	CreatedAt:           "device_template_status.created_at",
	UpdatedAt:           "device_template_status.updated_at",
}

// Generated where

var DeviceTemplateStatusWhere = struct {
	DeviceEthAddr       whereHelper__byte
	TemplateDBCURL      whereHelpernull_String
	TemplatePidURL      whereHelpernull_String
	TemplateSettingsURL whereHelpernull_String
	FirmwareVersion     whereHelpernull_String
	CreatedAt           whereHelpertime_Time
	UpdatedAt           whereHelpertime_Time
}{
	DeviceEthAddr:       whereHelper__byte{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"device_eth_addr\""},
	TemplateDBCURL:      whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"template_dbc_url\""},
	TemplatePidURL:      whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"template_pid_url\""},
	TemplateSettingsURL: whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"template_settings_url\""},
	FirmwareVersion:     whereHelpernull_String{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"firmware_version\""},
	CreatedAt:           whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"created_at\""},
	UpdatedAt:           whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"device_template_status\".\"updated_at\""},
}

// DeviceTemplateStatusRels is where relationship names are stored.
var DeviceTemplateStatusRels = struct {
}{}

// deviceTemplateStatusR is where relationships are stored.
type deviceTemplateStatusR struct {
}

// NewStruct creates a new relationship struct
func (*deviceTemplateStatusR) NewStruct() *deviceTemplateStatusR {
	return &deviceTemplateStatusR{}
}

// deviceTemplateStatusL is where Load methods for each relationship are stored.
type deviceTemplateStatusL struct{}

var (
	deviceTemplateStatusAllColumns            = []string{"device_eth_addr", "template_dbc_url", "template_pid_url", "template_settings_url", "firmware_version", "created_at", "updated_at"}
	deviceTemplateStatusColumnsWithoutDefault = []string{"device_eth_addr"}
	deviceTemplateStatusColumnsWithDefault    = []string{"template_dbc_url", "template_pid_url", "template_settings_url", "firmware_version", "created_at", "updated_at"}
	deviceTemplateStatusPrimaryKeyColumns     = []string{"device_eth_addr"}
	deviceTemplateStatusGeneratedColumns      = []string{}
)

type (
	// DeviceTemplateStatusSlice is an alias for a slice of pointers to DeviceTemplateStatus.
	// This should almost always be used instead of []DeviceTemplateStatus.
	DeviceTemplateStatusSlice []*DeviceTemplateStatus
	// DeviceTemplateStatusHook is the signature for custom DeviceTemplateStatus hook methods
	DeviceTemplateStatusHook func(context.Context, boil.ContextExecutor, *DeviceTemplateStatus) error

	deviceTemplateStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deviceTemplateStatusType                 = reflect.TypeOf(&DeviceTemplateStatus{})
	deviceTemplateStatusMapping              = queries.MakeStructMapping(deviceTemplateStatusType)
	deviceTemplateStatusPrimaryKeyMapping, _ = queries.BindMapping(deviceTemplateStatusType, deviceTemplateStatusMapping, deviceTemplateStatusPrimaryKeyColumns)
	deviceTemplateStatusInsertCacheMut       sync.RWMutex
	deviceTemplateStatusInsertCache          = make(map[string]insertCache)
	deviceTemplateStatusUpdateCacheMut       sync.RWMutex
	deviceTemplateStatusUpdateCache          = make(map[string]updateCache)
	deviceTemplateStatusUpsertCacheMut       sync.RWMutex
	deviceTemplateStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var deviceTemplateStatusAfterSelectMu sync.Mutex
var deviceTemplateStatusAfterSelectHooks []DeviceTemplateStatusHook

var deviceTemplateStatusBeforeInsertMu sync.Mutex
var deviceTemplateStatusBeforeInsertHooks []DeviceTemplateStatusHook
var deviceTemplateStatusAfterInsertMu sync.Mutex
var deviceTemplateStatusAfterInsertHooks []DeviceTemplateStatusHook

var deviceTemplateStatusBeforeUpdateMu sync.Mutex
var deviceTemplateStatusBeforeUpdateHooks []DeviceTemplateStatusHook
var deviceTemplateStatusAfterUpdateMu sync.Mutex
var deviceTemplateStatusAfterUpdateHooks []DeviceTemplateStatusHook

var deviceTemplateStatusBeforeDeleteMu sync.Mutex
var deviceTemplateStatusBeforeDeleteHooks []DeviceTemplateStatusHook
var deviceTemplateStatusAfterDeleteMu sync.Mutex
var deviceTemplateStatusAfterDeleteHooks []DeviceTemplateStatusHook

var deviceTemplateStatusBeforeUpsertMu sync.Mutex
var deviceTemplateStatusBeforeUpsertHooks []DeviceTemplateStatusHook
var deviceTemplateStatusAfterUpsertMu sync.Mutex
var deviceTemplateStatusAfterUpsertHooks []DeviceTemplateStatusHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DeviceTemplateStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DeviceTemplateStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DeviceTemplateStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DeviceTemplateStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DeviceTemplateStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DeviceTemplateStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DeviceTemplateStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DeviceTemplateStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DeviceTemplateStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deviceTemplateStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDeviceTemplateStatusHook registers your hook function for all future operations.
func AddDeviceTemplateStatusHook(hookPoint boil.HookPoint, deviceTemplateStatusHook DeviceTemplateStatusHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		deviceTemplateStatusAfterSelectMu.Lock()
		deviceTemplateStatusAfterSelectHooks = append(deviceTemplateStatusAfterSelectHooks, deviceTemplateStatusHook)
		deviceTemplateStatusAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		deviceTemplateStatusBeforeInsertMu.Lock()
		deviceTemplateStatusBeforeInsertHooks = append(deviceTemplateStatusBeforeInsertHooks, deviceTemplateStatusHook)
		deviceTemplateStatusBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		deviceTemplateStatusAfterInsertMu.Lock()
		deviceTemplateStatusAfterInsertHooks = append(deviceTemplateStatusAfterInsertHooks, deviceTemplateStatusHook)
		deviceTemplateStatusAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		deviceTemplateStatusBeforeUpdateMu.Lock()
		deviceTemplateStatusBeforeUpdateHooks = append(deviceTemplateStatusBeforeUpdateHooks, deviceTemplateStatusHook)
		deviceTemplateStatusBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		deviceTemplateStatusAfterUpdateMu.Lock()
		deviceTemplateStatusAfterUpdateHooks = append(deviceTemplateStatusAfterUpdateHooks, deviceTemplateStatusHook)
		deviceTemplateStatusAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		deviceTemplateStatusBeforeDeleteMu.Lock()
		deviceTemplateStatusBeforeDeleteHooks = append(deviceTemplateStatusBeforeDeleteHooks, deviceTemplateStatusHook)
		deviceTemplateStatusBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		deviceTemplateStatusAfterDeleteMu.Lock()
		deviceTemplateStatusAfterDeleteHooks = append(deviceTemplateStatusAfterDeleteHooks, deviceTemplateStatusHook)
		deviceTemplateStatusAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		deviceTemplateStatusBeforeUpsertMu.Lock()
		deviceTemplateStatusBeforeUpsertHooks = append(deviceTemplateStatusBeforeUpsertHooks, deviceTemplateStatusHook)
		deviceTemplateStatusBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		deviceTemplateStatusAfterUpsertMu.Lock()
		deviceTemplateStatusAfterUpsertHooks = append(deviceTemplateStatusAfterUpsertHooks, deviceTemplateStatusHook)
		deviceTemplateStatusAfterUpsertMu.Unlock()
	}
}

// One returns a single deviceTemplateStatus record from the query.
func (q deviceTemplateStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DeviceTemplateStatus, error) {
	o := &DeviceTemplateStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for device_template_status")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DeviceTemplateStatus records from the query.
func (q deviceTemplateStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeviceTemplateStatusSlice, error) {
	var o []*DeviceTemplateStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DeviceTemplateStatus slice")
	}

	if len(deviceTemplateStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DeviceTemplateStatus records in the query.
func (q deviceTemplateStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count device_template_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deviceTemplateStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if device_template_status exists")
	}

	return count > 0, nil
}

// DeviceTemplateStatuses retrieves all the records using an executor.
func DeviceTemplateStatuses(mods ...qm.QueryMod) deviceTemplateStatusQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"device_template_status\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"device_template_status\".*"})
	}

	return deviceTemplateStatusQuery{q}
}

// FindDeviceTemplateStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDeviceTemplateStatus(ctx context.Context, exec boil.ContextExecutor, deviceEthAddr []byte, selectCols ...string) (*DeviceTemplateStatus, error) {
	deviceTemplateStatusObj := &DeviceTemplateStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"device_template_status\" where \"device_eth_addr\"=$1", sel,
	)

	q := queries.Raw(query, deviceEthAddr)

	err := q.Bind(ctx, exec, deviceTemplateStatusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from device_template_status")
	}

	if err = deviceTemplateStatusObj.doAfterSelectHooks(ctx, exec); err != nil {
		return deviceTemplateStatusObj, err
	}

	return deviceTemplateStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DeviceTemplateStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no device_template_status provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(deviceTemplateStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deviceTemplateStatusInsertCacheMut.RLock()
	cache, cached := deviceTemplateStatusInsertCache[key]
	deviceTemplateStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deviceTemplateStatusAllColumns,
			deviceTemplateStatusColumnsWithDefault,
			deviceTemplateStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deviceTemplateStatusType, deviceTemplateStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deviceTemplateStatusType, deviceTemplateStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"device_template_status\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"device_template_status\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into device_template_status")
	}

	if !cached {
		deviceTemplateStatusInsertCacheMut.Lock()
		deviceTemplateStatusInsertCache[key] = cache
		deviceTemplateStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DeviceTemplateStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DeviceTemplateStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	deviceTemplateStatusUpdateCacheMut.RLock()
	cache, cached := deviceTemplateStatusUpdateCache[key]
	deviceTemplateStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deviceTemplateStatusAllColumns,
			deviceTemplateStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update device_template_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"device_template_status\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, deviceTemplateStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deviceTemplateStatusType, deviceTemplateStatusMapping, append(wl, deviceTemplateStatusPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update device_template_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for device_template_status")
	}

	if !cached {
		deviceTemplateStatusUpdateCacheMut.Lock()
		deviceTemplateStatusUpdateCache[key] = cache
		deviceTemplateStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q deviceTemplateStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for device_template_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for device_template_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DeviceTemplateStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceTemplateStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"device_template_status\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, deviceTemplateStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in deviceTemplateStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all deviceTemplateStatus")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DeviceTemplateStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no device_template_status provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(deviceTemplateStatusColumnsWithDefault, o)

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

	deviceTemplateStatusUpsertCacheMut.RLock()
	cache, cached := deviceTemplateStatusUpsertCache[key]
	deviceTemplateStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			deviceTemplateStatusAllColumns,
			deviceTemplateStatusColumnsWithDefault,
			deviceTemplateStatusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			deviceTemplateStatusAllColumns,
			deviceTemplateStatusPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert device_template_status, could not build update column list")
		}

		ret := strmangle.SetComplement(deviceTemplateStatusAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(deviceTemplateStatusPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert device_template_status, could not build conflict column list")
			}

			conflict = make([]string, len(deviceTemplateStatusPrimaryKeyColumns))
			copy(conflict, deviceTemplateStatusPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"device_template_status\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(deviceTemplateStatusType, deviceTemplateStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deviceTemplateStatusType, deviceTemplateStatusMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert device_template_status")
	}

	if !cached {
		deviceTemplateStatusUpsertCacheMut.Lock()
		deviceTemplateStatusUpsertCache[key] = cache
		deviceTemplateStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DeviceTemplateStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DeviceTemplateStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DeviceTemplateStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deviceTemplateStatusPrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"device_template_status\" WHERE \"device_eth_addr\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from device_template_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for device_template_status")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deviceTemplateStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no deviceTemplateStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from device_template_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for device_template_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DeviceTemplateStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(deviceTemplateStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceTemplateStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"device_template_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, deviceTemplateStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deviceTemplateStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for device_template_status")
	}

	if len(deviceTemplateStatusAfterDeleteHooks) != 0 {
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
func (o *DeviceTemplateStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDeviceTemplateStatus(ctx, exec, o.DeviceEthAddr)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DeviceTemplateStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeviceTemplateStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceTemplateStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"device_template_status\".* FROM \"vehicle_signal_decoding_api\".\"device_template_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, deviceTemplateStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DeviceTemplateStatusSlice")
	}

	*o = slice

	return nil
}

// DeviceTemplateStatusExists checks if the DeviceTemplateStatus row exists.
func DeviceTemplateStatusExists(ctx context.Context, exec boil.ContextExecutor, deviceEthAddr []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"device_template_status\" where \"device_eth_addr\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, deviceEthAddr)
	}
	row := exec.QueryRowContext(ctx, sql, deviceEthAddr)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if device_template_status exists")
	}

	return exists, nil
}

// Exists checks if the DeviceTemplateStatus row exists.
func (o *DeviceTemplateStatus) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DeviceTemplateStatusExists(ctx, exec, o.DeviceEthAddr)
}
