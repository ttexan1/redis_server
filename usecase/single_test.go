package usecase_test

import (
	"redis_app/command"
	"redis_app/domain"
	"redis_app/usecase"
	"testing"
)

type mockSingleStore struct{}

func (t *mockSingleStore) GetValue(key string) (domain.Single, *domain.Error) {
	if key == "existKey" {
		return domain.Single{Value: "string"}, nil
	}
	if key == "nilKey" {
		return domain.Single{}, domain.NewError(domain.RespErrorNilValue, nil)
	}
	return domain.Single{}, nil
}

func (t *mockSingleStore) SetValue(key string, value domain.Single) *domain.Error {
	return nil
}

func (t *mockSingleStore) Delete(key string) {
	return
}

func newMockSingle() usecase.Single {
	uc := usecase.NewUsecase(&mockSingleStore{}, nil)
	return uc.NewSingle()
}

func TestUsecaseGet(t *testing.T) {
	s := newMockSingle()
	if got := s.Get("existKey"); got != domain.RespBulkString("string") {
		t.Fatalf("Error Expected: %v, Got: %v", "string", got)
	}
}

func TestUsecaseSet(t *testing.T) {
	s := newMockSingle()
	type request struct {
		key   string
		value string
		args  []string
		want  domain.RespString
	}
	cases := map[string]request{
		"SetNxOK": {key: "nilKey", value: "value", args: []string{command.SetOptionNX}, want: domain.RespOK},
		"SetNxNG": {key: "existKey", value: "value", args: []string{command.SetOptionNX}, want: domain.RespErrorNilValue},
		"SetXxOK": {key: "existKey", value: "value", args: []string{command.SetOptionXX}, want: domain.RespOK},
		"SetXxNG": {key: "nilKey", value: "value", args: []string{command.SetOptionXX}, want: domain.RespErrorNilValue},
		"SetOK":   {key: "existKey", value: "value", args: []string{}, want: domain.RespOK},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if got := s.Set(c.key, c.value, c.args); got != c.want {
				t.Fatalf("Error Expected: %v, err: %v", c.want, got)
			}
		})
	}
}

func TestUsecaseIncrBy(t *testing.T) {
	s := newMockSingle()
	type request struct {
		key   string
		value string
		args  []string
		want  domain.RespString
	}
	cases := map[string]request{
		"IncrByOK":  {key: "nilKey", value: "10", args: []string{command.SetOptionNX}, want: domain.RespInteger(10)},
		"IncrByNG1": {key: "nilKey", value: "nonInt", args: []string{command.SetOptionNX}, want: domain.RespErrorWrongArgumentType},
		"IncrByNG2": {key: "existKey", value: "10", args: []string{command.SetOptionNX}, want: domain.RespErrorWrongArgumentType},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if got := s.IncrBy(c.key, c.value); got != c.want {
				t.Fatalf("Error Expected: %v, err: %v", c.want, got)
			}
		})
	}
}
