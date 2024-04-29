-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table aftermarket_device_to_template
    add constraint aftermarket_device_to_template_templates_template_name_fk
        foreign key (template_name) references templates;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table aftermarket_device_to_template drop constraint aftermarket_device_to_template_templates_template_name_fk;
-- +goose StatementEnd
