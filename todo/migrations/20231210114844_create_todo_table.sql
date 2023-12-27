-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todos (
    id UUID     PRIMARY KEY,
    created_by  SERIAL,
    assignee    SERIAL,
    description VARCHAR(255),
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos;
-- +goose StatementEnd
