-- name: CreateChallenge :one
INSERT INTO challenges (
    title,
    description,
    reward_amount,
    start_date,
    end_date,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetChallengeByID :one
SELECT * FROM challenges
WHERE id = $1;

-- name: ListActiveChallenges :many
SELECT * FROM challenges
WHERE status = 'active'
ORDER BY end_date ASC;

-- name: JoinChallenge :one
INSERT INTO user_challenges (
    user_id,
    challenge_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUserChallengeStatus :one
SELECT * FROM user_challenges
WHERE user_id = $1 AND challenge_id = $2;

-- name: UpdateUserChallengeStatus :one
UPDATE user_challenges
SET 
    status = $3,
    updated_at = NOW()
WHERE user_id = $1 AND challenge_id = $2
RETURNING *;
