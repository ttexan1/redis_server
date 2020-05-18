package command

// Simple Structure Command List
const (
	Get         = "get"
	Set         = "set"
	Del         = "del"
	IncrBy      = "incrby"
	Incr        = "incr"
	DecrBy      = "decrby"
	Decr        = "decr"
	IncrByFloat = "incrbyfloat"
	Append      = "append" // not implemented but can be by this application structure
	// and so on...
)

// SingleCommandList is the white list of simple key value data command
var SingleCommandList = map[string]struct{}{
	Get:         {},
	Set:         {},
	Del:         {},
	IncrBy:      {},
	Incr:        {},
	DecrBy:      {},
	Decr:        {},
	IncrByFloat: {},
	Append:      {},
}

// Options
const (
	SetOptionNX = "nx"
	SetOptionXX = "xx"
	SetOptionEX = "ex"
	SetOptionPX = "px"
)
