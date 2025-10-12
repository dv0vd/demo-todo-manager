package dto

import "time"

type NoteDTO struct {
	ID          uint64
	Title       string
	Description string
	UserId      uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
