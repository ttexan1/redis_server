package usecase

import (
	"redis_app/domain"
)

// Static is the interface for static command
type Static interface {
	Ping(string) domain.RespString
	Echo(string) domain.RespString
}

type static struct{}

// NewStatic returns the Static interface
func (uc *UseCase) NewStatic() Static {
	return &static{}
}

func (s *static) Ping(key string) domain.RespString {
	if key != "" {
		return domain.RespBulkString(key)
	}
	return domain.RespPong
}

func (s *static) Echo(arg string) domain.RespString {
	return domain.RespBulkString(arg)
}
