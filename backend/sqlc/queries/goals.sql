-- name: CreateGoal :one
INSERT INTO goals (
    user_id,
    title,
    description,
    target_amount,
    deadline
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetGoalByID :one
SELECT * FROM goals
WHERE id = $1;

-- name: ListGoalsByUserID :many
SELECT * FROM goals
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateGoalStatus :one
UPDATE goals
SET 
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateGoalAmount :one
UPDATE goals
SET 
    current_amount = current_amount + $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
