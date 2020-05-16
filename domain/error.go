package domain

// Error is the struct for error
type Error struct {
	FullMessage string
	Type        string
}

// ErrorTypes
const (
	ErrorTypeNilValue            = "$-1"
	ErrorTypeWrongArgumentNumber = "-ERR wrong number of arguments"
	ErrorTypeWrongArgumentType   = "-ERR wrong argument type"
	ErrorTypeWrongSyntax         = "-ERR wrong syntax"
	ErrorTypeUnknownCommand      = "-ERR unknown command"
	ErrorTypeIndexOutOfRange     = "-ERR index out of range"
	// and so on
)

func (e *Error) Error() string {
	return e.Type + ":" + e.FullMessage + "\r\n"
}
