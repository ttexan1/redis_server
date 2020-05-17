package command

// Static Command List
const (
	Ping = "ping"
	Echo = "echo"
	// and so on...
)

// StaticCommandList is the white list of static command
var StaticCommandList = map[string]struct{}{
	Ping: {},
	Echo: {},
}
