-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

-- template_vehicles
alter table template_vehicles
    drop constraint template_vehicles_template_name_fkey;

ALTER TABLE template_vehicles
    ADD CONSTRAINT template_vehicles_fk
        FOREIGN KEY (template_name) REFERENCES templates(template_name) ON UPDATE CASCADE;

-- template_device_definitions
alter table template_device_definitions
    drop constraint template_device_definitions_template_name_fkey;

ALTER TABLE template_device_definitions
    ADD CONSTRAINT template_device_definitions_fk
        FOREIGN KEY (template_name) REFERENCES templates(template_name) ON UPDATE CASCADE;

-- pid_configs
alter table pid_configs
    drop constraint pid_configs_template_name_fkey;

ALTER TABLE pid_configs
    ADD CONSTRAINT pid_configs_fk
        FOREIGN KEY (template_name) REFERENCES templates(template_name) ON UPDATE CASCADE;

-- device_settings
alter table device_settings
    drop constraint device_settings_template_name_fkey;

ALTER TABLE device_settings
    ADD CONSTRAINT device_settings_fk
        FOREIGN KEY (template_name) REFERENCES templates(template_name) ON UPDATE CASCADE;

-- dbc_files
alter table dbc_files
    drop constraint dbc_files_template_name_fkey;

ALTER TABLE dbc_files
    ADD CONSTRAINT dbc_files_fk
        FOREIGN KEY (template_name) REFERENCES templates(template_name) ON UPDATE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
-- we aint goin back
-- +goose StatementEnd
