-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table pid_configs add column signal_name varchar(50) not null;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table pid_configs drop column signal_name;
-- +goose StatementEnd
