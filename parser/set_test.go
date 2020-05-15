package parser

import (
	"testing"

	"redis_app/command"
	"redis_app/domain"
)

type testSetStore struct{}

func (ts testSetStore) Get(key string) (*domain.Single, error) {
	if key == "key" {
		return &domain.Single{Length: 5, Value: "value"}, nil
	}
	return &domain.Single{}, nil
}

func (ts testSetStore) Set(args []string) error {
	if low(args[2]) == "xx" {
		return nil
	}
	return nil
}

func TestSet(t *testing.T) {
	ts := testSetStore{}
	pr := &parser{
		Len:       3,
		Arguments: []string{"key", "value", "XX"},
		Directive: command.Get,
		Store:     ts,
	}
	result := pr.Set()
	if result != domain.ResponseOK {
		t.Fatal(result, "error")
	}
}
