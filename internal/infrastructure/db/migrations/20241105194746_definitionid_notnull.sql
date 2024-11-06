-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table template_device_definitions alter column definition_id set not null;
alter table template_device_definitions drop column device_definition_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table template_device_definitions alter column definition_id drop not null;
alter table template_device_definitions add column device_definition_id text not null default '';
-- +goose StatementEnd
