package command

// Static Command List
const (
	Ping = "ping"
	Echo = "echo"
	// and so on
)

// StaticCommandList is the white list for simple key value data
var StaticCommandList = map[string]struct{}{
	Ping: {},
	Echo: {},
}

// Simple Structure Command List
const (
	Get         = "get"
	Set         = "set"
	Del         = "del"
	IncrBy      = "incrby"
	IncrByFloat = "incrbyfloat"
	Append      = "append" // not implemented but can be by this application structure
	// and so on
)

// SingleCommandWhiteList is the white list for simple key value data
var SingleCommandWhiteList = map[string]struct{}{
	Get:         {},
	Set:         {},
	Del:         {},
	IncrBy:      {},
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
