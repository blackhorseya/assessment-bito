package model

import (
	"time"
)

// Pair is an entity that represents a pair.
type Pair struct {
	ID        string    `json:"id,omitempty"`
	Left      User      `json:"left,omitempty"`
	Right     User      `json:"right,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// NewPair is to create a new pair.
func NewPair(left User, right User) *Pair {
	return &Pair{
		ID:        "",
		Left:      left,
		Right:     right,
		CreatedAt: time.Time{},
	}
}
