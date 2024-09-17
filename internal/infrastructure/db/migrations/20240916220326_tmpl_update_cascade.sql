-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table template_vehicles
    drop constraint template_vehicles_template_name_fkey;

ALTER TABLE template_vehicles
    ADD CONSTRAINT template_vehicles_fk
        FOREIGN KEY (template_name) REFERENCES templates(template_name) ON UPDATE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
-- we aint goin back
-- +goose StatementEnd
