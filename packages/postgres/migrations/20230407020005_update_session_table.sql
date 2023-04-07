-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE sessions ALTER COLUMN session_token SET DATA TYPE text USING session_token::text;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE sessions ALTER COLUMN session_token SET DATA TYPE varchar(255) USING session_token::varchar(255);

-- +goose StatementEnd
