-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE pid_configs ADD COLUMN message_column_position INT;
UPDATE pid_configs SET message_column_position = 0 WHERE message_column_position IS NULL;
ALTER TABLE pid_configs ALTER COLUMN message_column_position SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

alter table pid_configs drop column message_column_position;
-- +goose StatementEnd
