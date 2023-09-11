-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

CREATE TABLE IF NOT EXISTS serial_to_template_overrides (
    serial text NOT NULL,
    template_name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT serial_to_template_overrides_pkey PRIMARY KEY (serial, template_name)
);
CREATE TABLE IF NOT EXISTS template_types (
    type_name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT template_types_pkey PRIMARY KEY (type_name)

);
CREATE TABLE IF NOT EXISTS templates (
    template_name text NOT NULL, 
    parent_template_name text DEFAULT NULL,
    template_type text REFERENCES template_types(type_name),
    version text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT templates_pkey PRIMARY KEY (template_name)
);
CREATE TABLE IF NOT EXISTS pid_configs (
    id BIGSERIAL,
    template_name text REFERENCES templates(template_name) NOT NULL,
    header byteA NOT NULL,
    mode byteA NOT NULL,
    pid byteA NOT NULL,
    formula text NOT NULL,
    interval_seconds INTEGER NOT NULL,
    protocol text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT pid_configs_pkey PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS device_settings (
    id BIGSERIAL,
    template_name text REFERENCES templates(template_name) NOT NULL,
    battery_critical_level_voltage text NOT NULL,
    safety_cut_out_voltage text NOT NULL,
    sleep_timer_event_driven_interval text NOT NULL,
    sleep_timer_event_driven_period text NOT NULL,
    sleep_timer_inactivity_after_sleep_interval text NOT NULL,
    sleep_timer_inactivity_fallback_interval text NOT NULL,
    wake_trigger_voltage_level text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT device_settings_pkey PRIMARY KEY (template_name)
);
CREATE TABLE IF NOT EXISTS template_vehicles (
    make_slug text,
    template_name text REFERENCES templates(template_name) NOT NULL,
    year_start INT,
    year_end INT,
    model_whitelist text[],
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT template_vehicles_pkey PRIMARY KEY (make_slug, year_start, year_end)
);

CREATE TABLE IF NOT EXISTS dbc_files (
    dbc_file text NOT NULL,
    template_name text REFERENCES templates(template_name) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT dbc_files_pkey PRIMARY KEY (template_name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE serial_to_template_overrides;
DROP TABLE pid_configs;
DROP TABLE device_settings;
DROP TABLE template_vehicles;
DROP TABLE dbc_files;
DROP TABLE templates;
DROP TABLE template_types;

-- +goose StatementEnd
