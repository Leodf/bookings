-- +goose Up
CREATE UNIQUE INDEX IF NOT EXISTS room_restrictions_start_date_end_date_idx ON room_restrictions(start_date, end_date);
CREATE UNIQUE INDEX IF NOT EXISTS room_restrictions_room_id_idx ON room_restrictions(room_id);
CREATE UNIQUE INDEX IF NOT EXISTS room_restrictions_reservation_id_idx ON room_restrictions(reservation_id);

-- +goose Down
DROP INDEX IF EXISTS room_restrictions_start_date_end_date_idx;
DROP INDEX IF EXISTS room_restrictions_room_id_idx;
DROP INDEX IF EXISTS room_restrictions_reservation_id_idx;
