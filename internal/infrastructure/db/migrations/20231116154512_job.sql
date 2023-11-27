-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

CREATE TYPE job_status_enum as enum ('PENDING', 'RUNNING', 'COMPLETED', 'FAILED');
CREATE TABLE IF NOT EXISTS jobs
(
    id character(27) COLLATE pg_catalog."default" NOT NULL,
    command text COLLATE pg_catalog."default" NOT NULL,
    status job_status_enum NOT NULL default 'PENDING',
    metadata jsonb,
    device_ethereum_address bytea NOT NULL,
    last_execution timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             CONSTRAINT job_pkey PRIMARY KEY (id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

SET search_path = vehicle_signal_decoding_api, public;
DROP TABLE jobs;

-- +goose StatementEnd
