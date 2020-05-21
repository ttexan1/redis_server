package main

import (
	"fmt"
	"net"
	"redis_app/domain"
	"testing"
	"time"
)

func TestStream(t *testing.T) {
	go main()
	time.Sleep(100 * time.Millisecond)
	const port = 10000
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	type testCase struct {
		req  string
		want domain.RespString
	}
	testCases := map[string]testCase{
		"ping":    {req: "*1\r\n$4\r\nping\r\n", want: domain.RespPong},
		"getFail": {req: "*2\r\n$3\r\nget\r\n$3\r\nkey\r\n", want: domain.RespErrorNilValue},
		"set":     {req: "*3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n", want: domain.RespOK},
		"getOK":   {req: "*2\r\n$3\r\nget\r\n$3\r\nkey\r\n", want: domain.RespBulkString("value")},
	}
	for key, tc := range testCases {
		t.Run(key, func(t *testing.T) {
			if _, err := conn.Write([]byte(tc.req)); err != nil {
				t.Fatalf("Connection Writing Error")
			}
			var got = make([]byte, 256)
			n, err := conn.Read(got)
			if err != nil {
				t.Fatalf("Connection Reading Error")
			}
			if string(got[:n]) != string(tc.want) {
				t.Fatalf("Error Expected: %v, Got: %v", string(tc.want), string(got))
			}
		})
	}
}
