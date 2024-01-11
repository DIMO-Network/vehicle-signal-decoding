-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

CREATE TABLE IF NOT EXISTS user_device_template
(
    device_ethereum_address bytea NOT NULL,
    template_dbc_url TEXT REFERENCES templates(template_name) NOT NULL,
    template_pid_url TEXT REFERENCES templates(template_name) NOT NULL,
    template_setting_url TEXT REFERENCES templates(template_name) NOT NULL,
    version text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_device_template_id_pkey PRIMARY KEY (device_ethereum_address)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;
DROP TABLE user_device_template;

-- +goose StatementEnd
