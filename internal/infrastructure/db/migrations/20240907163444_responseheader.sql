-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

ALTER TABLE pid_configs ADD COLUMN response_header bytea default '\x07e8'::bytea not null;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;

ALTER TABLE pid_configs DROP COLUMN response_header;
-- +goose StatementEnd
