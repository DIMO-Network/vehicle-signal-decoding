-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE device_settings
    DROP COLUMN IF EXISTS id,
    DROP COLUMN IF EXISTS battery_critical_level_voltage,
    DROP COLUMN IF EXISTS safety_cut_out_voltage,
    DROP COLUMN IF EXISTS sleep_timer_event_driven_interval,
    DROP COLUMN IF EXISTS sleep_timer_event_driven_period,
    DROP COLUMN IF EXISTS sleep_timer_inactivity_after_sleep_interval,
    DROP COLUMN IF EXISTS sleep_timer_inactivity_fallback_interval,
    DROP COLUMN IF EXISTS wake_trigger_voltage_level;

ALTER TABLE device_settings ADD COLUMN settings JSONB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE device_settings DROP COLUMN IF EXISTS settings;

ALTER TABLE device_settings
    ADD COLUMN id BIGSERIAL,
    ADD COLUMN battery_critical_level_voltage FLOAT,
    ADD COLUMN safety_cut_out_voltage FLOAT,
    ADD COLUMN sleep_timer_event_driven_interval FLOAT,
    ADD COLUMN sleep_timer_event_driven_period FLOAT,
    ADD COLUMN sleep_timer_inactivity_after_sleep_interval FLOAT,
    ADD COLUMN sleep_timer_inactivity_fallback_interval FLOAT,
    ADD COLUMN wake_trigger_voltage_level FLOAT;




-- +goose StatementEnd
