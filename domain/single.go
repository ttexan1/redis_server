package domain

import (
	"fmt"
	"strconv"
)

// Single is the sturct for the simple key value commands
type Single struct {
	// Key       string
	Value string

	// when you implement ex, px options, these column will be needed
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// ExpiredMs int64
}

// IntValue converts the value into int type
func (s *Single) IntValue() (int, *Error) {
	val, err := strconv.Atoi(s.Value)
	if err != nil {
		return val, &Error{Message: RespErrorWrongArgumentType}
	}
	return val, nil
}

// FloatValue converts the value into int type
func (s *Single) FloatValue() (float64, *Error) {
	val, err := strconv.ParseFloat(s.Value, 64)
	if err != nil {
		return 0, &Error{Message: RespErrorWrongArgumentType}
	}
	return val, nil
}

// RespString returns the response which satisfies the REdis Serialization Protocol
func (s *Single) RespString() RespString {
	return RespString(fmt.Sprintf("$%d\r\n%s\r\n", len(s.Value), s.Value))
}
