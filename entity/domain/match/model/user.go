package model

import (
	"time"
)

// User is an entity that represents a user.
type User struct {
	ID        string    `json:"id,omitempty"`
	Profile   Profile   `json:"profile,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
