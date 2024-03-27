-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table pid_configs drop column bytes_returned;
alter table pid_configs drop column message_column_position;
alter table pid_configs add column can_flow_control_clear boolean;
alter table pid_configs add column can_flow_control_id_pair varchar(20);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table pid_configs drop column can_flow_control_clear;
alter table pid_configs drop column can_flow_control_id_pair;
-- +goose StatementEnd
