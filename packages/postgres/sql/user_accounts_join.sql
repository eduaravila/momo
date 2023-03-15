
CREATE TABLE user_accounts_join (
    PRIMARY KEY (user_id, account_id),
    user_id uuid not null,
    account_id varchar(255) not null,
    account_manage_id varchar(255) not null,
    role_id SERIAL not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (account_id) REFERENCES accounts(id),    
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (account_manage_id) REFERENCES accounts(id)
);

-- name: CreateUserAccountsJoin :one
INSERT INTO user_accounts_join (user_id, account_id, account_manage_id, role_id) VALUES ($1, $2, $3, $4) on conflict (user_id, account_id) do nothing RETURNING *;

-- name: GetUserAccountsJoin :one
SELECT * FROM user_accounts_join WHERE user_id = $1 AND account_id = $2;

-- name: DeleteUserAccountsJoin :exec
DELETE FROM user_accounts_join WHERE user_id = $1 AND account_id = $2;

-- name: UpdateUserAccountsJoin :exec
UPDATE user_accounts_join SET updated_at = now() WHERE user_id = $1 AND account_id = $2;


-- name: ListUserAccountsJoins :many
SELECT * FROM user_accounts_join;

-- name: CountUserAccountsJoins :one
SELECT COUNT(*) FROM user_accounts_join;

-- name: GetUserAccountsJoinByUserID :many
SELECT * FROM user_accounts_join WHERE user_id = $1;

-- name: GetUserAccountsJoinByAccountID :many
SELECT * FROM user_accounts_join WHERE account_id = $1;

-- name: GetUserAccountsJoinByAccountManageID :many
SELECT * FROM user_accounts_join WHERE account_manage_id = $1;

-- name: GetUserAccountsJoinByRoleID :many
SELECT * FROM user_accounts_join WHERE role_id = $1;