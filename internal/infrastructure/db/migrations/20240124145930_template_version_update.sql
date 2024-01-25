-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

drop table if exists device_template_status;
CREATE TABLE IF NOT EXISTS device_template_status
(
    device_eth_addr bytea NOT NULL,
    -- all urls should be saved with the version
    template_dbc_url TEXT NULL,
    template_pid_url TEXT NULL,
    template_settings_url TEXT NULL,
    firmware_version TEXT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT device_template_status_id_pkey PRIMARY KEY (device_eth_addr)
);
-- add version to settings
alter table device_settings add column version text not null default 'v1.0.0';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;

drop table device_template_status;
alter table device_settings drop column version;

-- +goose StatementEnd
