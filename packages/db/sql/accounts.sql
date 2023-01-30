CREATE TABLE accounts (
    id varchar(255) primary key not null, 
    name varchar(255) not null,
    picture varchar(255) not null,
    prefered_username varchar(255) not null,
    access_token varchar(255) not null,
    refresh_token varchar(255) not null,
    platform_id varchar(255) not null,
    created_at timestamp not null default now(),
    expired_at timestamp not null,
    scope varchar(255) not null,
    FOREIGN KEY (platform_id) REFERENCES platforms(name)
);


-- name : CreateAccount :exec
INSERT INTO accounts (id, name, picture, prefered_username, access_token, refresh_token, platform_id, expired_at, scope) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) on conflict (id) do update set name = $2, picture = $3, prefered_username = $4, access_token = $5, refresh_token = $6, platform_id = $7, expired_at = $8, scope = $9;

-- name : GetAccount :one
SELECT * FROM accounts WHERE id = $1;

-- name : DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name : UpdateAccount :exec
UPDATE accounts SET updated_at = now() WHERE id = $1;

-- name : ListAccounts :many
SELECT * FROM accounts;

-- name : CountAccounts :one
SELECT COUNT(*) FROM accounts;

-- name : GetAccountsByIDs :many
SELECT * FROM accounts WHERE id = ANY($1);

-- name : GetAccountsByPlatform :many
SELECT * FROM accounts WHERE platform_id = $1;

-- name : GetAccountsByPlatformAndIDs :many
SELECT * FROM accounts WHERE platform_id = $1 AND id = ANY($2);

-- name : GetAccountsByPlatformAndName :many
SELECT * FROM accounts WHERE platform_id = $1 AND name = $2;

-- name : GetAccountsByPlatformAndNameAndIDs :many
SELECT * FROM accounts WHERE platform_id = $1 AND name = $2 AND id = ANY($3);
