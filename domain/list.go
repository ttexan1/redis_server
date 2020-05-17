package domain

import "time"

// ListValue is the struct for list type command
type ListValue struct {
	Value     []string
	CreatedAt time.Time
}
