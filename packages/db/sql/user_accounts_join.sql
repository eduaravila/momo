
CREATE TABLE user_accounts_join (
    PRIMARY KEY (user_id, account_id),
    user_id uuid not null,
    account_id varchar(255) not null,
    role_id SERIAL not null,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);
-- name: CreateUserAccountsJoin :exec
INSERT INTO user_accounts_join (user_id, account_id, role_id) VALUES ($1, $2, $3) on conflict (user_id, account_id) do update set role_id = $3;

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
SELECT * FROM user_accounts_join INNER JOIN accounts ON user_accounts_join.account_id = accounts.id WHERE user_id = $1;

-- name: GetAccountByAccountID :one
SELECT * FROM accounts INNER JOIN user_accounts_join ON accounts.id = user_accounts_join.account_id WHERE account_id = $1;

-- name: UserAccountsJoinExistByAccountID :one
SELECT EXISTS (SELECT 1 FROM user_accounts_join WHERE account_id = $1);
