-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;
-- keeps track of device to last template requested mapping
CREATE TABLE IF NOT EXISTS device_template_status
(
    vin char(17) NOT NULL,
    device_eth_addr bytea NULL,
    template_dbc_url TEXT REFERENCES templates(template_name) NOT NULL,
    template_pid_url TEXT REFERENCES templates(template_name) NOT NULL,
    template_setting_url TEXT REFERENCES templates(template_name) NOT NULL,
    template_version text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_device_template_id_pkey PRIMARY KEY (vin)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;
DROP TABLE device_template;

-- +goose StatementEnd
