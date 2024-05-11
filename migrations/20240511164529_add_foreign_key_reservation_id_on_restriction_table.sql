-- +goose Up
ALTER TABLE room_restrictions
ADD CONSTRAINT room_restrictions_reservation_id_fk
FOREIGN KEY (reservation_id) REFERENCES reservations(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +goose Down
ALTER TABLE room_restrictions
DROP CONSTRAINT room_restrictions_reservation_id_fk;
