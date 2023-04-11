-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table test_signals add signals jsonb;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table test_signals drop column signals;
-- +goose StatementEnd
