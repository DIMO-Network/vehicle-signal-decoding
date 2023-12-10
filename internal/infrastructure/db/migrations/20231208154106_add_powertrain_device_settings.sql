-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SET search_path = vehicle_signal_decoding_api, public;

ALTER TABLE device_settings
    ADD COLUMN powertrain varchar(4) NOT NULL default 'ICE';

ALTER TABLE device_settings
    ADD CONSTRAINT check_allowed_values CHECK (powertrain IN ('ICE', 'HEV', 'PHEV', 'BEV', 'FCEV'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE device_settings
    DROP COLUMN IF EXISTS powertrain;
-- +goose StatementEnd
