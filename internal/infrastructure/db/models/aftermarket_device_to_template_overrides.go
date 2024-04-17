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

// AftermarketDeviceToTemplateOverride is an object representing the database table.
type AftermarketDeviceToTemplateOverride struct {
	AftermarketDeviceEthereumAddress []byte    `boil:"aftermarket_device_ethereum_address" json:"aftermarket_device_ethereum_address" toml:"aftermarket_device_ethereum_address" yaml:"aftermarket_device_ethereum_address"`
	TemplateName                     string    `boil:"template_name" json:"template_name" toml:"template_name" yaml:"template_name"`
	CreatedAt                        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                        time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *aftermarketDeviceToTemplateOverrideR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L aftermarketDeviceToTemplateOverrideL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AftermarketDeviceToTemplateOverrideColumns = struct {
	AftermarketDeviceEthereumAddress string
	TemplateName                     string
	CreatedAt                        string
	UpdatedAt                        string
}{
	AftermarketDeviceEthereumAddress: "aftermarket_device_ethereum_address",
	TemplateName:                     "template_name",
	CreatedAt:                        "created_at",
	UpdatedAt:                        "updated_at",
}

var AftermarketDeviceToTemplateOverrideTableColumns = struct {
	AftermarketDeviceEthereumAddress string
	TemplateName                     string
	CreatedAt                        string
	UpdatedAt                        string
}{
	AftermarketDeviceEthereumAddress: "aftermarket_device_to_template_overrides.aftermarket_device_ethereum_address",
	TemplateName:                     "aftermarket_device_to_template_overrides.template_name",
	CreatedAt:                        "aftermarket_device_to_template_overrides.created_at",
	UpdatedAt:                        "aftermarket_device_to_template_overrides.updated_at",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AftermarketDeviceToTemplateOverrideWhere = struct {
	AftermarketDeviceEthereumAddress whereHelper__byte
	TemplateName                     whereHelperstring
	CreatedAt                        whereHelpertime_Time
	UpdatedAt                        whereHelpertime_Time
}{
	AftermarketDeviceEthereumAddress: whereHelper__byte{field: "\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\".\"aftermarket_device_ethereum_address\""},
	TemplateName:                     whereHelperstring{field: "\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\".\"template_name\""},
	CreatedAt:                        whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\".\"created_at\""},
	UpdatedAt:                        whereHelpertime_Time{field: "\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\".\"updated_at\""},
}

// AftermarketDeviceToTemplateOverrideRels is where relationship names are stored.
var AftermarketDeviceToTemplateOverrideRels = struct {
}{}

// aftermarketDeviceToTemplateOverrideR is where relationships are stored.
type aftermarketDeviceToTemplateOverrideR struct {
}

// NewStruct creates a new relationship struct
func (*aftermarketDeviceToTemplateOverrideR) NewStruct() *aftermarketDeviceToTemplateOverrideR {
	return &aftermarketDeviceToTemplateOverrideR{}
}

// aftermarketDeviceToTemplateOverrideL is where Load methods for each relationship are stored.
type aftermarketDeviceToTemplateOverrideL struct{}

var (
	aftermarketDeviceToTemplateOverrideAllColumns            = []string{"aftermarket_device_ethereum_address", "template_name", "created_at", "updated_at"}
	aftermarketDeviceToTemplateOverrideColumnsWithoutDefault = []string{"aftermarket_device_ethereum_address", "template_name"}
	aftermarketDeviceToTemplateOverrideColumnsWithDefault    = []string{"created_at", "updated_at"}
	aftermarketDeviceToTemplateOverridePrimaryKeyColumns     = []string{"aftermarket_device_ethereum_address", "template_name"}
	aftermarketDeviceToTemplateOverrideGeneratedColumns      = []string{}
)

type (
	// AftermarketDeviceToTemplateOverrideSlice is an alias for a slice of pointers to AftermarketDeviceToTemplateOverride.
	// This should almost always be used instead of []AftermarketDeviceToTemplateOverride.
	AftermarketDeviceToTemplateOverrideSlice []*AftermarketDeviceToTemplateOverride
	// AftermarketDeviceToTemplateOverrideHook is the signature for custom AftermarketDeviceToTemplateOverride hook methods
	AftermarketDeviceToTemplateOverrideHook func(context.Context, boil.ContextExecutor, *AftermarketDeviceToTemplateOverride) error

	aftermarketDeviceToTemplateOverrideQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	aftermarketDeviceToTemplateOverrideType                 = reflect.TypeOf(&AftermarketDeviceToTemplateOverride{})
	aftermarketDeviceToTemplateOverrideMapping              = queries.MakeStructMapping(aftermarketDeviceToTemplateOverrideType)
	aftermarketDeviceToTemplateOverridePrimaryKeyMapping, _ = queries.BindMapping(aftermarketDeviceToTemplateOverrideType, aftermarketDeviceToTemplateOverrideMapping, aftermarketDeviceToTemplateOverridePrimaryKeyColumns)
	aftermarketDeviceToTemplateOverrideInsertCacheMut       sync.RWMutex
	aftermarketDeviceToTemplateOverrideInsertCache          = make(map[string]insertCache)
	aftermarketDeviceToTemplateOverrideUpdateCacheMut       sync.RWMutex
	aftermarketDeviceToTemplateOverrideUpdateCache          = make(map[string]updateCache)
	aftermarketDeviceToTemplateOverrideUpsertCacheMut       sync.RWMutex
	aftermarketDeviceToTemplateOverrideUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var aftermarketDeviceToTemplateOverrideAfterSelectMu sync.Mutex
var aftermarketDeviceToTemplateOverrideAfterSelectHooks []AftermarketDeviceToTemplateOverrideHook

var aftermarketDeviceToTemplateOverrideBeforeInsertMu sync.Mutex
var aftermarketDeviceToTemplateOverrideBeforeInsertHooks []AftermarketDeviceToTemplateOverrideHook
var aftermarketDeviceToTemplateOverrideAfterInsertMu sync.Mutex
var aftermarketDeviceToTemplateOverrideAfterInsertHooks []AftermarketDeviceToTemplateOverrideHook

var aftermarketDeviceToTemplateOverrideBeforeUpdateMu sync.Mutex
var aftermarketDeviceToTemplateOverrideBeforeUpdateHooks []AftermarketDeviceToTemplateOverrideHook
var aftermarketDeviceToTemplateOverrideAfterUpdateMu sync.Mutex
var aftermarketDeviceToTemplateOverrideAfterUpdateHooks []AftermarketDeviceToTemplateOverrideHook

var aftermarketDeviceToTemplateOverrideBeforeDeleteMu sync.Mutex
var aftermarketDeviceToTemplateOverrideBeforeDeleteHooks []AftermarketDeviceToTemplateOverrideHook
var aftermarketDeviceToTemplateOverrideAfterDeleteMu sync.Mutex
var aftermarketDeviceToTemplateOverrideAfterDeleteHooks []AftermarketDeviceToTemplateOverrideHook

var aftermarketDeviceToTemplateOverrideBeforeUpsertMu sync.Mutex
var aftermarketDeviceToTemplateOverrideBeforeUpsertHooks []AftermarketDeviceToTemplateOverrideHook
var aftermarketDeviceToTemplateOverrideAfterUpsertMu sync.Mutex
var aftermarketDeviceToTemplateOverrideAfterUpsertHooks []AftermarketDeviceToTemplateOverrideHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AftermarketDeviceToTemplateOverride) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AftermarketDeviceToTemplateOverride) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AftermarketDeviceToTemplateOverride) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AftermarketDeviceToTemplateOverride) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AftermarketDeviceToTemplateOverride) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AftermarketDeviceToTemplateOverride) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AftermarketDeviceToTemplateOverride) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AftermarketDeviceToTemplateOverride) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AftermarketDeviceToTemplateOverride) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aftermarketDeviceToTemplateOverrideAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAftermarketDeviceToTemplateOverrideHook registers your hook function for all future operations.
func AddAftermarketDeviceToTemplateOverrideHook(hookPoint boil.HookPoint, aftermarketDeviceToTemplateOverrideHook AftermarketDeviceToTemplateOverrideHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		aftermarketDeviceToTemplateOverrideAfterSelectMu.Lock()
		aftermarketDeviceToTemplateOverrideAfterSelectHooks = append(aftermarketDeviceToTemplateOverrideAfterSelectHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		aftermarketDeviceToTemplateOverrideBeforeInsertMu.Lock()
		aftermarketDeviceToTemplateOverrideBeforeInsertHooks = append(aftermarketDeviceToTemplateOverrideBeforeInsertHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		aftermarketDeviceToTemplateOverrideAfterInsertMu.Lock()
		aftermarketDeviceToTemplateOverrideAfterInsertHooks = append(aftermarketDeviceToTemplateOverrideAfterInsertHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		aftermarketDeviceToTemplateOverrideBeforeUpdateMu.Lock()
		aftermarketDeviceToTemplateOverrideBeforeUpdateHooks = append(aftermarketDeviceToTemplateOverrideBeforeUpdateHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		aftermarketDeviceToTemplateOverrideAfterUpdateMu.Lock()
		aftermarketDeviceToTemplateOverrideAfterUpdateHooks = append(aftermarketDeviceToTemplateOverrideAfterUpdateHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		aftermarketDeviceToTemplateOverrideBeforeDeleteMu.Lock()
		aftermarketDeviceToTemplateOverrideBeforeDeleteHooks = append(aftermarketDeviceToTemplateOverrideBeforeDeleteHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		aftermarketDeviceToTemplateOverrideAfterDeleteMu.Lock()
		aftermarketDeviceToTemplateOverrideAfterDeleteHooks = append(aftermarketDeviceToTemplateOverrideAfterDeleteHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		aftermarketDeviceToTemplateOverrideBeforeUpsertMu.Lock()
		aftermarketDeviceToTemplateOverrideBeforeUpsertHooks = append(aftermarketDeviceToTemplateOverrideBeforeUpsertHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		aftermarketDeviceToTemplateOverrideAfterUpsertMu.Lock()
		aftermarketDeviceToTemplateOverrideAfterUpsertHooks = append(aftermarketDeviceToTemplateOverrideAfterUpsertHooks, aftermarketDeviceToTemplateOverrideHook)
		aftermarketDeviceToTemplateOverrideAfterUpsertMu.Unlock()
	}
}

// One returns a single aftermarketDeviceToTemplateOverride record from the query.
func (q aftermarketDeviceToTemplateOverrideQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AftermarketDeviceToTemplateOverride, error) {
	o := &AftermarketDeviceToTemplateOverride{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for aftermarket_device_to_template_overrides")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AftermarketDeviceToTemplateOverride records from the query.
func (q aftermarketDeviceToTemplateOverrideQuery) All(ctx context.Context, exec boil.ContextExecutor) (AftermarketDeviceToTemplateOverrideSlice, error) {
	var o []*AftermarketDeviceToTemplateOverride

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AftermarketDeviceToTemplateOverride slice")
	}

	if len(aftermarketDeviceToTemplateOverrideAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AftermarketDeviceToTemplateOverride records in the query.
func (q aftermarketDeviceToTemplateOverrideQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count aftermarket_device_to_template_overrides rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q aftermarketDeviceToTemplateOverrideQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if aftermarket_device_to_template_overrides exists")
	}

	return count > 0, nil
}

// AftermarketDeviceToTemplateOverrides retrieves all the records using an executor.
func AftermarketDeviceToTemplateOverrides(mods ...qm.QueryMod) aftermarketDeviceToTemplateOverrideQuery {
	mods = append(mods, qm.From("\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\".*"})
	}

	return aftermarketDeviceToTemplateOverrideQuery{q}
}

// FindAftermarketDeviceToTemplateOverride retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAftermarketDeviceToTemplateOverride(ctx context.Context, exec boil.ContextExecutor, aftermarketDeviceEthereumAddress []byte, templateName string, selectCols ...string) (*AftermarketDeviceToTemplateOverride, error) {
	aftermarketDeviceToTemplateOverrideObj := &AftermarketDeviceToTemplateOverride{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" where \"aftermarket_device_ethereum_address\"=$1 AND \"template_name\"=$2", sel,
	)

	q := queries.Raw(query, aftermarketDeviceEthereumAddress, templateName)

	err := q.Bind(ctx, exec, aftermarketDeviceToTemplateOverrideObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from aftermarket_device_to_template_overrides")
	}

	if err = aftermarketDeviceToTemplateOverrideObj.doAfterSelectHooks(ctx, exec); err != nil {
		return aftermarketDeviceToTemplateOverrideObj, err
	}

	return aftermarketDeviceToTemplateOverrideObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AftermarketDeviceToTemplateOverride) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no aftermarket_device_to_template_overrides provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(aftermarketDeviceToTemplateOverrideColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	aftermarketDeviceToTemplateOverrideInsertCacheMut.RLock()
	cache, cached := aftermarketDeviceToTemplateOverrideInsertCache[key]
	aftermarketDeviceToTemplateOverrideInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			aftermarketDeviceToTemplateOverrideAllColumns,
			aftermarketDeviceToTemplateOverrideColumnsWithDefault,
			aftermarketDeviceToTemplateOverrideColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(aftermarketDeviceToTemplateOverrideType, aftermarketDeviceToTemplateOverrideMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(aftermarketDeviceToTemplateOverrideType, aftermarketDeviceToTemplateOverrideMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into aftermarket_device_to_template_overrides")
	}

	if !cached {
		aftermarketDeviceToTemplateOverrideInsertCacheMut.Lock()
		aftermarketDeviceToTemplateOverrideInsertCache[key] = cache
		aftermarketDeviceToTemplateOverrideInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AftermarketDeviceToTemplateOverride.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AftermarketDeviceToTemplateOverride) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	aftermarketDeviceToTemplateOverrideUpdateCacheMut.RLock()
	cache, cached := aftermarketDeviceToTemplateOverrideUpdateCache[key]
	aftermarketDeviceToTemplateOverrideUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			aftermarketDeviceToTemplateOverrideAllColumns,
			aftermarketDeviceToTemplateOverridePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update aftermarket_device_to_template_overrides, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, aftermarketDeviceToTemplateOverridePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(aftermarketDeviceToTemplateOverrideType, aftermarketDeviceToTemplateOverrideMapping, append(wl, aftermarketDeviceToTemplateOverridePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update aftermarket_device_to_template_overrides row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for aftermarket_device_to_template_overrides")
	}

	if !cached {
		aftermarketDeviceToTemplateOverrideUpdateCacheMut.Lock()
		aftermarketDeviceToTemplateOverrideUpdateCache[key] = cache
		aftermarketDeviceToTemplateOverrideUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q aftermarketDeviceToTemplateOverrideQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for aftermarket_device_to_template_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for aftermarket_device_to_template_overrides")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AftermarketDeviceToTemplateOverrideSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aftermarketDeviceToTemplateOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, aftermarketDeviceToTemplateOverridePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in aftermarketDeviceToTemplateOverride slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all aftermarketDeviceToTemplateOverride")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AftermarketDeviceToTemplateOverride) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no aftermarket_device_to_template_overrides provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(aftermarketDeviceToTemplateOverrideColumnsWithDefault, o)

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

	aftermarketDeviceToTemplateOverrideUpsertCacheMut.RLock()
	cache, cached := aftermarketDeviceToTemplateOverrideUpsertCache[key]
	aftermarketDeviceToTemplateOverrideUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			aftermarketDeviceToTemplateOverrideAllColumns,
			aftermarketDeviceToTemplateOverrideColumnsWithDefault,
			aftermarketDeviceToTemplateOverrideColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			aftermarketDeviceToTemplateOverrideAllColumns,
			aftermarketDeviceToTemplateOverridePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert aftermarket_device_to_template_overrides, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(aftermarketDeviceToTemplateOverridePrimaryKeyColumns))
			copy(conflict, aftermarketDeviceToTemplateOverridePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(aftermarketDeviceToTemplateOverrideType, aftermarketDeviceToTemplateOverrideMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(aftermarketDeviceToTemplateOverrideType, aftermarketDeviceToTemplateOverrideMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert aftermarket_device_to_template_overrides")
	}

	if !cached {
		aftermarketDeviceToTemplateOverrideUpsertCacheMut.Lock()
		aftermarketDeviceToTemplateOverrideUpsertCache[key] = cache
		aftermarketDeviceToTemplateOverrideUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AftermarketDeviceToTemplateOverride record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AftermarketDeviceToTemplateOverride) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AftermarketDeviceToTemplateOverride provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), aftermarketDeviceToTemplateOverridePrimaryKeyMapping)
	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" WHERE \"aftermarket_device_ethereum_address\"=$1 AND \"template_name\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from aftermarket_device_to_template_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for aftermarket_device_to_template_overrides")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q aftermarketDeviceToTemplateOverrideQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no aftermarketDeviceToTemplateOverrideQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from aftermarket_device_to_template_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for aftermarket_device_to_template_overrides")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AftermarketDeviceToTemplateOverrideSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(aftermarketDeviceToTemplateOverrideBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aftermarketDeviceToTemplateOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, aftermarketDeviceToTemplateOverridePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from aftermarketDeviceToTemplateOverride slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for aftermarket_device_to_template_overrides")
	}

	if len(aftermarketDeviceToTemplateOverrideAfterDeleteHooks) != 0 {
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
func (o *AftermarketDeviceToTemplateOverride) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAftermarketDeviceToTemplateOverride(ctx, exec, o.AftermarketDeviceEthereumAddress, o.TemplateName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AftermarketDeviceToTemplateOverrideSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AftermarketDeviceToTemplateOverrideSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aftermarketDeviceToTemplateOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\".* FROM \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, aftermarketDeviceToTemplateOverridePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AftermarketDeviceToTemplateOverrideSlice")
	}

	*o = slice

	return nil
}

// AftermarketDeviceToTemplateOverrideExists checks if the AftermarketDeviceToTemplateOverride row exists.
func AftermarketDeviceToTemplateOverrideExists(ctx context.Context, exec boil.ContextExecutor, aftermarketDeviceEthereumAddress []byte, templateName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vehicle_signal_decoding_api\".\"aftermarket_device_to_template_overrides\" where \"aftermarket_device_ethereum_address\"=$1 AND \"template_name\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, aftermarketDeviceEthereumAddress, templateName)
	}
	row := exec.QueryRowContext(ctx, sql, aftermarketDeviceEthereumAddress, templateName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if aftermarket_device_to_template_overrides exists")
	}

	return exists, nil
}

// Exists checks if the AftermarketDeviceToTemplateOverride row exists.
func (o *AftermarketDeviceToTemplateOverride) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AftermarketDeviceToTemplateOverrideExists(ctx, exec, o.AftermarketDeviceEthereumAddress, o.TemplateName)
}
