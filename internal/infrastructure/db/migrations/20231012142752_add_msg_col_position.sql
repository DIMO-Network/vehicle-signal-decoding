-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE pid_configs ADD COLUMN message_column_position INT DEFAULT 0;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

alter table pid_configs drop column message_column_position;
-- +goose StatementEnd
