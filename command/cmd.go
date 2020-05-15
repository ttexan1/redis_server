package command

// Static Command List
const (
	Ping = "ping"
	Echo = "echo"
	// and so on
)

// Simple Structure Command List
const (
	Get    = "get"
	Set    = "set"
	Del    = "del"
	IncrBy = "incrby"
	Append = "append" // not implemented but can be by this application structure
	// and so on
)

// List Structure Command List
const (
	Lpush  = "lpush"
	Lrange = "lrange"
	// and so on
)

// System Command List
// Memory Command List
// and so on

// Options
const (
	SetOptionNX = "nx"
	SetOptionXX = "xx"
	SetOptionEX = "ex"
	SetOptionPX = "px"
)

// Typical error responses
const (
	InvalidFormat = "invalid_format"
)
