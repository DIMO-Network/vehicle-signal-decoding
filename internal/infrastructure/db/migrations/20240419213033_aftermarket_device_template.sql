-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;
DROP TABLE IF EXISTS serial_to_template_overrides CASCADE;
CREATE TABLE IF NOT EXISTS aftermarket_device_to_template (
                                                              aftermarket_device_ethereum_address bytea NOT NULL,
                                                              template_name TEXT NOT NULL,
                                                              created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                              updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

                                                              CONSTRAINT aftermarket_device_to_template_pkey PRIMARY KEY (aftermarket_device_ethereum_address, template_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
CREATE TABLE IF NOT EXISTS serial_to_template_overrides (
                                                            serial TEXT NOT NULL,
                                                            template_name TEXT NOT NULL,
                                                            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

                                                            CONSTRAINT serial_to_template_overrides_pkey PRIMARY KEY (serial, template_name)
);
DROP TABLE IF EXISTS aftermarket_device_to_template CASCADE;
-- +goose StatementEnd
