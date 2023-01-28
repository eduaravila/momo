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
    platform uuid not null,
    access_token varchar(255) not null,
    refresh_token varchar(255) not null,
    created_at timestamp not null,
    expired_at timestamp not null,
    scope varchar(255) not null,
    sub varchar(255) not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (platform) REFERENCES platforms(id)
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

CREATE TABLE platforms (
    id uuid primary key not null,
    name varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE playlists (
    id uuid primary key not null,
    user_id uuid not null,
    platform uuid not null,
    name varchar(255) not null,
    description varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (platform) REFERENCES platforms(id)
);

CREATE TABLE playlist_tracks (
    id uuid primary key not null,
    playlist_id uuid not null,
    track_id uuid not null,
    FOREIGN KEY (playlist_id) REFERENCES playlists(id),
    FOREIGN KEY (track_id) REFERENCES tracks(id)
);

CREATE TABLE tracks (
    id uuid primary key not null,
    title varchar(255) not null,
    artist varchar(255) not null,
    album varchar(255) not null,
    album_cover varchar(255) not null,
    duration int not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE tracks_history (
    id uuid primary key not null,
    user_id uuid not null,
    track_id uuid not null,
    created_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (track_id) REFERENCES tracks(id)
);

CREATE TABLE playlists_history (
    id uuid primary key not null,
    user_id uuid not null,
    playlist_id uuid not null,
    created_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (playlist_id) REFERENCES playlists(id)
);

CREATE TABLE tracks_history (
    id uuid primary key not null,
    user_id uuid not null,
    track_id uuid not null,
    created_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (track_id) REFERENCES tracks(id)
);

CREATE TABLE playlists_history (
    id uuid primary key not

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd
