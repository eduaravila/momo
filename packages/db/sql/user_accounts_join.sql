
-- name : CreateUserAccountsJoin :exec
INSERT INTO user_accounts_join (user_id, account_id, role_id) VALUES ($1, $2, $3) on conflict (user_id, account_id) do update set role_id = $3;

-- name : GetUserAccountsJoin :one
SELECT * FROM user_accounts_join WHERE user_id = $1 AND account_id = $2;

-- name : DeleteUserAccountsJoin :exec
DELETE FROM user_accounts_join WHERE user_id = $1 AND account_id = $2;

-- name : UpdateUserAccountsJoin :exec
UPDATE user_accounts_join SET updated_at = now() WHERE user_id = $1 AND account_id = $2;

-- name : ListUserAccountsJoins :many
SELECT * FROM user_accounts_join;

-- name : CountUserAccountsJoins :one
SELECT COUNT(*) FROM user_accounts_join;


-- name : GetUserAccountsJoinByUserID :many
SELECT * FROM user_accounts_join INNER JOIN accounts ON user_accounts_join.account_id = accounts.id WHERE user_id = $1;
