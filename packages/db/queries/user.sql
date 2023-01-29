
-- name: CreateUser :exec 
INSERT INTO
    users (
        id,
        name,
        created_at,
        picture,
        prefered_username,
        updated_at
    )
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: UpdateUser :exec
UPDATE
    users
SET
    name = $2,
    created_at = $3,
    picture = $4,
    prefered_username = $5,
    updated_at = $6
WHERE
    id = $1;

-- name: GetUsers :many
SELECT
    id,
    name,
    created_at,
    picture,
    prefered_username,
    updated_at
FROM
    users
ORDER BY
    created_at DESC;

-- name: GetUser :one
SELECT
    id,
    name,
    created_at,
    picture,
    prefered_username,
    updated_at
FROM
    users
WHERE
    id = $1;

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = $1;