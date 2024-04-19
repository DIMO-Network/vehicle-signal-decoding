-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
ALTER TABLE pid_configs ALTER COLUMN protocol DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
ALTER TABLE pid_configs ALTER COLUMN protocol set NOT NULL;
-- +goose StatementEnd
