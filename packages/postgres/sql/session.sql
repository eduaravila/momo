
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



-- name: CreateSession :one
INSERT INTO sessions (id, user_id, created_at, expired_at, session_token, ip_address, user_agent, is_valid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;

-- name: InvalidateSession :exec
UPDATE sessions SET is_valid = false WHERE id = $1;




