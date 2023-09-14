-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

alter table templates add column protocol can_protocol_type NOT NULL;
alter table templates add column powertrain text NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

alter table templates drop column protocol;
alter table templates drop column powertrain;

-- +goose StatementEnd
