-- +goose Up
INSERT INTO rooms (room_name, created_at, updated_at) VALUES ('General''s Quarters', now(), now());
INSERT INTO rooms (room_name, created_at, updated_at) VALUES ('Major''s Suite', now(), now());

-- +goose Down
DELETE FROM rooms;

