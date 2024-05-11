-- +goose Up
CREATE TABLE IF NOT EXISTS restrictions (
    id SERIAL PRIMARY KEY,
    restriction_name VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE restrictions;
