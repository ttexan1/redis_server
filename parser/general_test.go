package parser

import (
	"testing"

	"redis_app/command"
	"redis_app/domain"
	"redis_app/store"
)

const pingSimple = `*1
$4
ping
`

const pingWithArgument = `*2
$4
ping
$3
aaa
`

func TestRawParser(t *testing.T) {
	pr := rawStringToArguments(pingWithArgument)
	if pr.Len != 2 {
		t.Fatal("Wrong Length")
	}
	if pr.Directive != command.Ping {
		t.Fatal("Wrong Command")
	}
	if pr.Arguments[0] != "aaa" {
		t.Fatal("Wrong Argument")
	}
}

func TestPingCommand(t *testing.T) {
	res := ParseCommand(pingSimple, &store.DB{})
	if res != domain.ResponsePong {
		t.Fatal(res, "Wrong Result")
	}
}

const setArgument1 = `*3
$3
set
$3
key
$5
value
`
const setArgument2 = `*4
$3
set
$4
key2
$6
value2
$2
xx
`

func TestSetGeneral(t *testing.T) {
	db := &store.DB{St: make(map[string]*domain.Single)}
	res := ParseCommand(setArgument1, db)
	if res != domain.ResponseOK {
		t.Fatal(res, "Wrong Result")
	}
	res = ParseCommand(setArgument2, db)
}
