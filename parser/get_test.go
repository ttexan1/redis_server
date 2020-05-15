package parser

import (
	"testing"

	"redis_app/command"
	"redis_app/domain"
)

type testGetStore struct{}

func (ts testGetStore) Get(key string) (*domain.Single, error) {
	if key == "key" {
		return &domain.Single{Length: 5, Value: "value"}, nil
	}
	return &domain.Single{}, nil
}

func (ts testGetStore) Set(args []string) error {
	return nil
}

func TestGet(t *testing.T) {
	ts := testGetStore{}
	pr := &parser{
		Len:       2,
		Arguments: []string{"key"},
		Directive: command.Get,
		Store:     ts,
	}
	result := pr.Get()
	if result != "$5\nvalue\r\n" {
		t.Fatal(result, "error")
	}
}
