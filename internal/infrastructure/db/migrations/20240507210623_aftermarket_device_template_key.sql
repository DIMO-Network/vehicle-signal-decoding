-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table aftermarket_device_to_template
    drop constraint aftermarket_device_to_template_pkey;

alter table aftermarket_device_to_template add constraint aftermarket_device_to_template_pkey primary key (aftermarket_device_ethereum_address);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table aftermarket_device_to_template drop constraint aftermarket_device_to_template_pkey;

alter table aftermarket_device_to_template add constraint aftermarket_device_to_template_pkey primary key (aftermarket_device_ethereum_address, template_name);
-- +goose StatementEnd
