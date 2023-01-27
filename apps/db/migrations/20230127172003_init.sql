-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid primary key not null,
    name varchar(255) not null,
    created_at timestamp not null,
    picture varchar(255) not null,
    prefered_usrname varchar(255) not null,
    updated_at timestamp not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd
