-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE device_settings
    ADD COLUMN powertrain TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE device_settings
    DROP COLUMN IF EXISTS powertrain;
-- +goose StatementEnd
