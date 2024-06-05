-- +goose Up
DROP INDEX IF EXISTS room_restrictions_start_date_end_date_idx;
CREATE INDEX IF NOT EXISTS room_restrictions_start_date_end_date_idx ON room_restrictions(start_date, end_date);

-- +goose Down
DROP INDEX IF EXISTS room_restrictions_start_date_end_date_idx;
CREATE UNIQUE INDEX IF NOT EXISTS room_restrictions_start_date_end_date_idx ON room_restrictions(start_date, end_date);
