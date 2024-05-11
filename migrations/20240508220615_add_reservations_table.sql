-- +goose Up
CREATE TABLE IF NOT EXISTS reservations (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(255),
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    room_id INTEGER
);

-- +goose Down
DROP TABLE reservations;
