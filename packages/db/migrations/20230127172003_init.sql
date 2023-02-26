-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid primary key not null,   
    created_at timestamp not null default now(),
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

CREATE TABLE issuers (
    name varchar(255) primary key not null,
    is_active boolean not null default true,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);


CREATE TABLE accounts (
    id uuid primary key not null,
    user_id uuid not null,
    picture varchar(255) not null,
    email varchar(255) not null,
    prefered_username varchar(255) not null,
    access_token varchar(255) not null,
    refresh_token varchar(255) not null,
    iss varchar(255) not null,
    sub varchar(255) not null unique,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    expired_at timestamp not null,
    scope varchar(255) not null,
    FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE role_permissions (
    role_id INTEGER REFERENCES roles(id),
    permission_id INTEGER REFERENCES permissions(id),
    PRIMARY KEY (role_id, permission_id)
);

-- save mods and other roles here an user can be in multiple roles, have multiple permissions, and have multiple accounts
CREATE TABLE user_accounts_join (
    PRIMARY KEY (user_id, account_id),
    user_id uuid not null,
    account_id uuid not null,
    account_manage_id uuid not null,
    role_id SERIAL not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (account_id) REFERENCES accounts(id),    
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (account_manage_id) REFERENCES accounts(id)
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

DROP TABLE issuers cascade;

DROP TABLE files;

DROP TABLE roles cascade;

DROP TABLE permissions cascade;

DROP TABLE role_permissions cascade;

DROP TABLE user_accounts_join cascade;

DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd