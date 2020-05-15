package parser

import (
	"testing"

	"redis_app/domain"
)

func TestPing(t *testing.T) {
	pr := &parser{
		Directive: "ping",
		Len:       1,
		Arguments: []string{},
	}
	if pr.Ping() != domain.ResponsePong {
		t.Fatal(pr.Ping(), "Not Correct")
	}
}
