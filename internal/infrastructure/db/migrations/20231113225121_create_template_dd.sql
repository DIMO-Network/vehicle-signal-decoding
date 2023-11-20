-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS template_device_definitions (
    id BIGSERIAL,
    device_definition_id char(27) NOT NULL,
    device_style_id char(27),
    template_name TEXT REFERENCES templates(template_name) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT template_device_definitions_pkey PRIMARY KEY (id),
    CONSTRAINT device_def_unique UNIQUE (device_definition_id, template_name, device_style_id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS template_device_definitions;

-- +goose StatementEnd
