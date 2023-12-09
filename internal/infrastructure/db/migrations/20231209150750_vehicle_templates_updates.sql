-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table template_vehicles
    drop constraint template_vehicles_pkey;

drop index template_vehicles_pkey;

alter table template_vehicles
    alter column make_slug drop not null;

alter table template_vehicles
    add id serial;

alter table template_vehicles
    add primary key (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;

alter table template_vehicles
    alter column make_slug set not null;

alter table template_vehicles
    drop constraint template_vehicles_pkey;

alter table template_vehicles
    add primary key (make_slug, year_start, year_end);

alter table template_vehicles
    drop column id;

-- +goose StatementEnd
