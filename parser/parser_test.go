package parser

import (
	"testing"
)

func TestRawParser(t *testing.T) {

	arguments := map[string]*respRequest{
		"*1\r\n$4\r\nping\r\n":                              {Directive: "ping", Len: 1},
		"*2\r\n$4\r\nping\r\n$3\r\nabc\r\n":                 {Directive: "ping", Len: 2},
		"*3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n":   {Directive: "set", Len: 3},
		"*3\r\n$3\r\nset\r\n$4\r\nkey2\r\n$6\r\nvalue2\r\n": {Directive: "set", Len: 3},
	}

	for key, value := range arguments {
		t.Run(key, func(t *testing.T) {
			if ps, err := decode(key); err != nil ||
				(ps.Directive != value.Directive) ||
				(ps.Len != value.Len) {
				t.Errorf("Expect: %#v, Got: %#v", value, ps)
			}
		})
	}
}
