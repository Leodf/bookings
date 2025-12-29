-- +goose Up
-- +goose StatementBegin
ALTER TABLE reservations
ADD COLUMN processed INTEGER DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE reservations
DROP COLUMN IF EXISTS processed;
-- +goose StatementEnd
