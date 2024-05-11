-- +goose Up
CREATE TABLE IF NOT EXISTS room_restrictions (
    id SERIAL PRIMARY KEY,
    start_date DATE,
    end_date DATE,
    room_id INTEGER,
    reservation_id INTEGER,
    restriction_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE room_restrictions;
