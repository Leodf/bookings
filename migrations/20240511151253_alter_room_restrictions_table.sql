-- +goose Up
ALTER TABLE room_restrictions
ADD CONSTRAINT room_restrictions_rooms_id_fk
FOREIGN KEY (room_id) REFERENCES rooms(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE room_restrictions
ADD CONSTRAINT room_restrictions_restrictions_id_fk
FOREIGN KEY (restriction_id) REFERENCES restrictions(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +goose Down
ALTER TABLE room_restrictions
DROP CONSTRAINT room_restrictions_rooms_id_fk;

ALTER TABLE room_restrictions
DROP CONSTRAINT room_restrictions_restrictions_id_fk;

