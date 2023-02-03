CREATE TABLE accounts (
    id uuid primary key not null,
    name varchar(255) not null,
    picture varchar(255) not null,
    email varchar(255) not null,
    sub varchar(255) not null unique,
    prefered_username varchar(255) not null,
    access_token varchar(255) not null,
    refresh_token varchar(255) not null,
    iss varchar(255) not null,
    created_at timestamp not null default now(),
    expired_at timestamp not null,
    scope varchar(255) not null,
    FOREIGN KEY (iss) REFERENCES issuers(name)
);

-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    name,
    picture,
    email,
    sub,
    prefered_username,
    access_token,
    refresh_token,
    iss,
    expired_at,
    scope
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
)ON CONFLICT (sub) DO UPDATE SET
    name = $2,
    picture = $3,
    email = $4,
    prefered_username = $6,
    access_token = $7,
    refresh_token = $8,
    iss = $9,
    expired_at = $10,
    scope = $11 
    RETURNING id;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: UpdateAccount :exec
UPDATE accounts SET updated_at = now() WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM accounts;

-- name: CountAccounts :one
SELECT COUNT(*) FROM accounts;

-- name: GetAccountsByIDs :many
SELECT * FROM accounts WHERE id = ANY($1);

-- name: GetAccountsByISS :many
SELECT * FROM accounts WHERE iss = $1;

-- name: AccountExistByEmail :one
SELECT EXISTS (SELECT 1 FROM accounts WHERE email = $1);

-- name: AccountExistBySub :one
SELECT EXISTS (SELECT 1 FROM accounts WHERE sub = $1);

