package domain

import "fmt"

// RespString is the string which satisfies the redis serialization protocol
type RespString string

// RespInteger returns RespInteger
func RespInteger(arg int) RespString {
	return RespString(fmt.Sprintf(":%d\r\n", arg))
}

// RespErr returns RespErr
func RespErr(arg string) RespString {
	return RespString(ErrorPrefix + ": " + arg + "\r\n")
}

// RespBulkString returns RespBulkString
func RespBulkString(arg string) RespString {
	str := fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	return RespString(str)
}

// RespArray returns RespArray
func RespArray(arg string) RespString {
	return ""
}

// Response for success
const (
	RespPong = "+PONG\r\n"
	RespOK   = "+OK\r\n"
)
