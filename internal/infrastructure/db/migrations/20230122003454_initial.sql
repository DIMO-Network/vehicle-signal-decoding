-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'trigger_enum') THEN
    CREATE TYPE trigger_enum AS ENUM ('CAN', 'PID');
  END IF;
END $$;


CREATE TABLE IF NOT EXISTS dbc_codes
(
    id character(27) COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    dbc_contents text COLLATE pg_catalog."default" NULL,
    header integer NULL,
    trigger trigger_enum NOT NULL default 'CAN',
    recording_enabled boolean NOT NULL default true,
    max_sample_size integer NOT NULL default 5, -- how often do we want to record per autopi for this signal
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
    autopi_unit_id char(38) not null,
    value text COLLATE pg_catalog."default" NOT NULL,
    approved boolean not null,
    vehicle_timestamp timestamp without time zone NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT test_signals_pkey PRIMARY KEY (id),
    CONSTRAINT fk_dbc_codes FOREIGN KEY (dbc_codes_id) REFERENCES dbc_codes (id)
);
DO $$
BEGIN
   IF NOT EXISTS (
       SELECT 1 
       FROM   pg_class c
       JOIN   pg_namespace n ON n.oid = c.relnamespace
       WHERE  c.relname = 'idx_autopi_unit_id' 
       AND    n.nspname = 'vehicle_signal_decoding_api' 
   ) THEN
       CREATE INDEX idx_autopi_unit_id ON test_signals(autopi_unit_id);
   END IF;
END $$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;
DROP TABLE test_signals;
DROP TABLE dbc_codes;

-- +goose StatementEnd
