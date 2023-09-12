-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table pid_configs add column bytes_returned smallint;

CREATE TYPE can_protocol_type AS ENUM (
    'CAN11_500',
    'CAN29_500'
    );

alter table pid_configs alter column header set default E'\\x07DF';
alter table pid_configs alter column mode set default E'\\x01';
alter table pid_configs alter column protocol TYPE can_protocol_type using protocol::can_protocol_type;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table pid_configs drop column bytes_returned;
alter table pid_configs alter column protocol TYPE varchar(50);
drop type can_protocol_type;

-- +goose StatementEnd
