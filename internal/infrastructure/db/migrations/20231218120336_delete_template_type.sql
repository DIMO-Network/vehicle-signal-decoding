-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE templates DROP COLUMN IF EXISTS template_type;
DROP TABLE IF EXISTS template_type;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

CREATE TABLE template_type (
    type_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE templates ADD COLUMN template_type VARCHAR(255);
-- +goose StatementEnd
