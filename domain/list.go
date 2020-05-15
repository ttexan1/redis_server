package domain

import "time"

// ListValue is the struct for list command type key-value
type ListValue struct {
	Value     []string
	CreatedAt time.Time
}
