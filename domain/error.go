package domain

// Error is the struct for error
type Error struct {
	Message RespString
	Err     error // wrapping sub error
}

// ErrorTypes
const (
	ErrorPrefix = "-ERR"

	RespErrorNilValue            = "$-1\r\n"
	RespErrorWrongArgumentNumber = "-ERR wrong number of arguments\r\n"
	RespErrorWrongArgumentType   = "-ERR wrong argument type\r\n"
	RespErrorWrongSyntax         = "-ERR wrong syntax\r\n"
	RespErrorUnknownCommand      = "-ERR unknown command\r\n"
	RespErrorIndexOutOfRange     = "-ERR index out of range\r\n"
	RespErrorWrongOperation      = "-ERR WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	// and so on...
)

// NewError returns the new error
func NewError(msg RespString, err error) *Error {
	return &Error{
		Message: msg,
		Err:     err,
	}
}

// RespError returns the error string which satisfies Redis Serialization Protocol
func (e *Error) RespError() RespString {
	if e.Err != nil {
		return RespErr(e.Err.Error())
	}
	return e.Message
}
