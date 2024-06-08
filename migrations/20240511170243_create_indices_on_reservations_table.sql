-- +goose Up
CREATE INDEX IF NOT EXISTS reservations_email_idx ON reservations(email);
CREATE INDEX IF NOT EXISTS reservations_last_name_idx ON reservations(last_name);

-- +goose Down
DROP INDEX IF EXISTS reservations_email_idx;
DROP INDEX IF EXISTS reservations_last_name_idx;
