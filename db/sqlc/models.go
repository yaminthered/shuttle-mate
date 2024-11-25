// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          int64          `json:"id"`
	GroupID     int64          `json:"group_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	StartTime   time.Time      `json:"start_time"`
	EndTime     sql.NullTime   `json:"end_time"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type EventParticipant struct {
	ID       int64        `json:"id"`
	EventID  int64        `json:"event_id"`
	UserID   int64        `json:"user_id"`
	JoinedAt sql.NullTime `json:"joined_at"`
}

type Group struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Location    sql.NullString `json:"location"`
	CreatedBy   int64          `json:"created_by"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type GroupMember struct {
	ID       int64          `json:"id"`
	GroupID  int64          `json:"group_id"`
	UserID   int64          `json:"user_id"`
	Role     sql.NullString `json:"role"`
	JoinedAt sql.NullTime   `json:"joined_at"`
}

type Message struct {
	ID        int64        `json:"id"`
	GroupID   int64        `json:"group_id"`
	UserID    int64        `json:"user_id"`
	Content   string       `json:"content"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type User struct {
	ID          int64          `json:"id"`
	PhoneNumber string         `json:"phone_number"`
	Name        sql.NullString `json:"name"`
	AvatarUrl   sql.NullString `json:"avatar_url"`
	Bio         sql.NullString `json:"bio"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}
