package dto

import "time"

type NoteDTO struct {
	ID          uint64
	Title       string
	Description string
	Done        bool
	UserId      uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
