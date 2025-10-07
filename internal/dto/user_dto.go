package dto

import "time"

type UserDTO struct {
	ID        uint64
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
