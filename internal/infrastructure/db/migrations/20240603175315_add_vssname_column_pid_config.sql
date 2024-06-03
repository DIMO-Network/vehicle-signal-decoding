-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
ALTER TABLE pid_configs ADD COLUMN vss_covesa_name TEXT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
ALTER TABLE pid_configs DROP COLUMN vss_covesa_name;
-- +goose StatementEnd