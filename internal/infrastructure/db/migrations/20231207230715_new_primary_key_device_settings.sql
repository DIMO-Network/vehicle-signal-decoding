-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

ALTER TABLE device_settings
    DROP CONSTRAINT IF EXISTS device_settings_pkey;

ALTER TABLE device_settings
    ADD COLUMN name TEXT default 'default-ice' PRIMARY KEY;

ALTER TABLE device_settings
    ALTER COLUMN template_name DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE device_settings
    ALTER COLUMN template_name SET NOT NULL,
    ADD CONSTRAINT device_settings_pkey PRIMARY KEY (template_name);

ALTER TABLE device_settings
DROP COLUMN IF EXISTS name;


-- +goose StatementEnd
