package domain

import "time"

// Single is the sturct for the simple key value commands
type Single struct {
	Value     string
	Length    int
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredMs uint64
}
