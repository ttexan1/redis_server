package domain

import (
	"strconv"
)

// Single is the sturct for the simple key value commands
type Single struct {
	// Key       string
	Value  string
	Length int

	// when you implement ex, px options, these column will be needed
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// ExpiredMs int64
}

// IntValue converts the value into int type
func (s *Single) IntValue() (int, error) {
	return strconv.Atoi(s.Value)
}

// FloatValue converts the value into int type
func (s *Single) FloatValue() (float64, error) {
	return strconv.ParseFloat(s.Value, 64)
}
