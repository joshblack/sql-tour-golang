package model

import (
	"time"
)

type Post struct {
	Title      string    `db:"title"`
	UserId     int64     `db:"user_id"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
