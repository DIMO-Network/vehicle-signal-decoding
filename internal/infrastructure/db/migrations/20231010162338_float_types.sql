-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

alter table device_settings alter column battery_critical_level_voltage type float USING NULLIF(battery_critical_level_voltage, '')::FLOAT;
alter table device_settings alter column safety_cut_out_voltage type float USING NULLIF(safety_cut_out_voltage, '')::FLOAT;
alter table device_settings alter column sleep_timer_event_driven_interval type float USING NULLIF(sleep_timer_event_driven_interval, '')::FLOAT;
alter table device_settings alter column sleep_timer_event_driven_period type float USING NULLIF(sleep_timer_event_driven_period, '')::FLOAT;
alter table device_settings alter column sleep_timer_inactivity_after_sleep_interval type float USING NULLIF(sleep_timer_inactivity_after_sleep_interval, '')::FLOAT;
alter table device_settings alter column sleep_timer_inactivity_fallback_interval type float USING NULLIF(sleep_timer_inactivity_fallback_interval, '')::FLOAT;
alter table device_settings alter column wake_trigger_voltage_level type float USING NULLIF(wake_trigger_voltage_level, '')::FLOAT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;

alter table device_settings alter column battery_critical_level_voltage type text;
alter table device_settings alter column safety_cut_out_voltage type text;
alter table device_settings alter column sleep_timer_event_driven_interval type text;
alter table device_settings alter column sleep_timer_event_driven_period type text;
alter table device_settings alter column sleep_timer_inactivity_after_sleep_interval type text;
alter table device_settings alter column sleep_timer_inactivity_fallback_interval type text;
alter table device_settings alter column wake_trigger_voltage_level type text;
-- +goose StatementEnd
