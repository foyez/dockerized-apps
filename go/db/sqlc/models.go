// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}
