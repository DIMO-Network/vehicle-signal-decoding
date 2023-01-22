-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

CREATE TABLE IF NOT EXISTS dbc_codes
(
    id character(27) COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    dbc_contents text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT dbc_codes_pkey PRIMARY KEY (id),
    CONSTRAINT dbc_codes_name_key UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS test_signals
(
    id character(27) COLLATE pg_catalog."default" NOT NULL,
    device_definition_id char(27) not null,
    dbc_codes_id char(27) not null,
    user_device_id char(27) not null,
    trigger text COLLATE pg_catalog."default" NOT NULL,
    signal_name text COLLATE pg_catalog."default" NOT NULL,
    value text COLLATE pg_catalog."default" NOT NULL,
    validated boolean not null,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT test_signals_pkey PRIMARY KEY (id),
    CONSTRAINT fk_dbc_codes FOREIGN KEY (dbc_codes_id) REFERENCES dbc_codes (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;
DROP TABLE test_signals;
DROP TABLE dbc_codes;

-- +goose StatementEnd
