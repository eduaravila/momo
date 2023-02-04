-- name: CreateUser :one 
INSERT INTO users (id) VALUES ($1) on conflict (id) do nothing RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET updated_at = now() WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: GetUsersByIDs :many
SELECT * FROM users WHERE id = ANY($1);
