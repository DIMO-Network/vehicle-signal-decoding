-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter type can_protocol_type ADD Value 'CAN11_250';
alter type can_protocol_type ADD Value 'CAN29_250';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
-- not possible to drop enum values
-- +goose StatementEnd
