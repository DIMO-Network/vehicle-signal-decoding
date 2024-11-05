-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table template_device_definitions add column definition_id text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table template_device_definitions drop column definition_id;
-- +goose StatementEnd
