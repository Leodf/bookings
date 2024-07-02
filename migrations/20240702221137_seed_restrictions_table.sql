-- +goose Up
INSERT INTO restrictions (restriction_name, created_at, updated_at) VALUES ('Reservation', now(), now());
INSERT INTO restrictions (restriction_name, created_at, updated_at) VALUES ('Owner block', now(), now());

-- +goose Down
DELETE FROM restrictions;
