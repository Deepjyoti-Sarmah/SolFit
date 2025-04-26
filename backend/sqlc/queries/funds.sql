-- name: CreateFund :one
INSERT INTO funds (
    user_id,
    goal_id,
    amount,
    transaction_hash,
    status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetFundByID :one
SELECT * FROM funds
WHERE id = $1;

-- name: ListFundsByGoalID :many
SELECT * FROM funds
WHERE goal_id = $1
ORDER BY created_at DESC;

-- name: UpdateFundStatus :one
UPDATE funds
SET 
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
