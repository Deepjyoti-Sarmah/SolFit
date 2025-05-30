// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Challenge struct {
	ID           int32              `json:"id"`
	Title        string             `json:"title"`
	Description  pgtype.Text        `json:"description"`
	RewardAmount pgtype.Numeric     `json:"reward_amount"`
	StartDate    pgtype.Timestamptz `json:"start_date"`
	EndDate      pgtype.Timestamptz `json:"end_date"`
	Status       string             `json:"status"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

type Fund struct {
	ID              int32              `json:"id"`
	UserID          int32              `json:"user_id"`
	GoalID          int32              `json:"goal_id"`
	Amount          pgtype.Numeric     `json:"amount"`
	TransactionHash pgtype.Text        `json:"transaction_hash"`
	Status          string             `json:"status"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}

type Goal struct {
	ID            int32              `json:"id"`
	UserID        int32              `json:"user_id"`
	Title         string             `json:"title"`
	Description   pgtype.Text        `json:"description"`
	TargetAmount  pgtype.Numeric     `json:"target_amount"`
	CurrentAmount pgtype.Numeric     `json:"current_amount"`
	Deadline      pgtype.Timestamptz `json:"deadline"`
	Status        string             `json:"status"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type Task struct {
	ID          int32              `json:"id"`
	GoalID      int32              `json:"goal_id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	Status      string             `json:"status"`
	DueDate     pgtype.Timestamptz `json:"due_date"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	ID            int32              `json:"id"`
	Username      string             `json:"username"`
	Email         string             `json:"email"`
	PasswordHash  string             `json:"password_hash"`
	WalletAddress pgtype.Text        `json:"wallet_address"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type UserChallenge struct {
	ID          int32              `json:"id"`
	UserID      int32              `json:"user_id"`
	ChallengeID int32              `json:"challenge_id"`
	Status      string             `json:"status"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
