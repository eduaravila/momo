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

CREATE TABLE accounts (
    id uuid primary key not null,
    user_id uuid not null,
    access_token varchar(255) not null,
    refresh_token varchar(255) not null,
    created_at timestamp not null,
    expired_at timestamp not null,
    scope varchar(255) not null,
    sub varchar(255) not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
);

CREATE TABLE sessions (
    id uuid primary key not null,
    user_id uuid not null,
    created_at timestamp not null,
    expired_at timestamp not null,
    session_token varchar(255) not null,
    ip_address varchar(255) not null,
    user_agent varchar(255) not null,
    is_valid boolean not null default true,
    FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE overlays (
    id uuid primary key not null,
    is_banned boolean not null default false,
    is_enabled boolean not null default true,
    user_id uuid not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE messages (
    id uuid primary key not null,
    user_id uuid not null,
    message text not null,
    created_at timestamp not null,
    audio_url varchar(255) not null,
    author_name varchar(255) not null,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE configurations (
    id uuid primary key not null,
    user_id uuid not null,
    max_message_length integer not null,
    min_amount_tip integer not null,
    min_point_tip integer not null,
    fallback_voice uuid not null,
    is_channel_point_enabled boolean not null,
    is_tipping_enabled boolean not null,
    is_bits_enabled boolean not null,
    is_sub_enabled boolean not null,
    is_donation_enabled boolean not null,
    is_command_enabled boolean not null,
    is_chat_enabled boolean not null,
    is_following_enabled boolean not null,
    is_hosting_enabled boolean not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (fallback_voice) REFERENCES voices(id)
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE accounts;
DROP TABLE sessions;
DROP TABLE overlays;
DROP TABLE messages;
DROP TABLE configurations;
DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd
