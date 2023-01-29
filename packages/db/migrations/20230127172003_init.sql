-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid primary key not null,
    name varchar(255) not null,
    created_at timestamp not null default now(),
    picture varchar(255) not null,
    prefered_username varchar(255) not null,
    updated_at timestamp not null default now()
);

CREATE TABLE voices (
    id uuid primary key not null,
    name varchar(255) not null,
    created_at timestamp not null default now()
);

CREATE TABLE files (
    id uuid primary key not null,
    file_name varchar(255) not null,
    file_path varchar(255) not null,
    file_type varchar(255) not null,
    created_at timestamp not null default now()
);

CREATE TABLE platforms (
    name varchar(255) primary key not null,
    is_active boolean not null default true,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

CREATE TABLE accounts (
    id uuid primary key not null,
    user_id uuid not null,
    access_token varchar(255) not null,
    refresh_token varchar(255) not null,
    platform_id varchar(255) not null,
    created_at timestamp not null default now(),
    expired_at timestamp not null,
    scope varchar(255) not null,
    sub varchar(255) not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (platform_id) REFERENCES platforms(name)
);

CREATE TABLE sessions (
    id uuid primary key not null,
    user_id uuid not null,
    created_at timestamp not null default now(),
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
    account_id uuid not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE messages (
    id uuid primary key not null,
    account_id uuid not null,
    message text not null,
    created_at timestamp not null default now(),
    audio_file_id uuid,
    author_name varchar(255) not null,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (audio_file_id) REFERENCES files(id)
);

CREATE TABLE configurations (
    id uuid primary key not null,
    account_id uuid not null,
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
    banned_word_file_id uuid not null,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (fallback_voice) REFERENCES voices(id),
    FOREIGN KEY (banned_word_file_id) REFERENCES files(id)
);

CREATE TABLE configurations_voices_banned_join (
    PRIMARY KEY (configuration_id, voice_id),
    configuration_id uuid not null,
    voice_id uuid not null,
    FOREIGN KEY (configuration_id) REFERENCES configurations(id),
    FOREIGN KEY (voice_id) REFERENCES voices(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users cascade;

DROP TABLE accounts cascade;

DROP TABLE sessions cascade;

DROP TABLE overlays;

DROP TABLE messages;

DROP TABLE configurations cascade;

DROP TABLE voices cascade;

DROP TABLE configurations_voices_banned_join;

DROP TABLE platforms;

DROP TABLE files;

DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd