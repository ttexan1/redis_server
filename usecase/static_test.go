package usecase_test

import (
	// "redis_app/command"
	"redis_app/domain"
	"redis_app/usecase"
	"testing"
)

func TestUsecasePingEcho(t *testing.T) {
	uc := usecase.NewUsecase(nil, nil)
	s := uc.NewStatic()
	if got := s.Ping(""); got != domain.RespPong {
		t.Fatalf("Error Expected: %v, Got: %v", domain.RespPong, got)
	}
	if got := s.Ping("ping"); got != domain.RespBulkString("ping") {
		t.Fatalf("Error Expected: %v, Got: %v", domain.RespBulkString("ping"), got)
	}
	if got := s.Echo("echo"); got != domain.RespBulkString("echo") {
		t.Fatalf("Error Expected: %v, Got: %v", domain.RespBulkString("echo"), got)
	}
}
