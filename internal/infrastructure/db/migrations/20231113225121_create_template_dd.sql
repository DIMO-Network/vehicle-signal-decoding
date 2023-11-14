-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS template_device_definitions (
    id BIGSERIAL,
    device_definition_id BIGINT NOT NULL,
    template_name TEXT REFERENCES templates(template_name) NOT NULL,
    year_start INTEGER NOT NULL,
    year_end INTEGER NOT NULL,
    make_slug TEXT NOT NULL,
    model_whitelist TEXT[] NOT NULL, -- array of model slugs, corresponding to ModelWhitelist in TemplateVehicle
    hierarchy_level INTEGER NOT NULL, -- Determines the specificity level in the hierarchy
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT template_device_definitions_pkey PRIMARY KEY (id),
    CONSTRAINT device_def_unique UNIQUE (device_definition_id, template_name, year_start, year_end, make_slug, hierarchy_level) -- Ensures uniqueness for each combination
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS template_device_definitions;

-- +goose StatementEnd
