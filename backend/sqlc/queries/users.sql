-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password_hash
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: UpdateUserWalletAddress :one 
UPDATE users 
SET 
  wallet_address = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;
