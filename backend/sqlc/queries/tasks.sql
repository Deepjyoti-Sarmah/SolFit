-- name: CreateTask :one
INSERT INTO tasks (
    goal_id,
    title,
    description,
    due_date
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;

-- name: ListTasksByGoalID :many
SELECT * FROM tasks
WHERE goal_id = $1
ORDER BY due_date ASC NULLS LAST;

-- name: UpdateTaskStatus :one
UPDATE tasks
SET 
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
