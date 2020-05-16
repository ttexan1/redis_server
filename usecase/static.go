package usecase

import (
	"fmt"
	"redis_app/domain"
)

// Static is the interface for static command
type Static interface {
	Ping(string) string
	Echo(string) string
}

type static struct{}

func (uc *UseCase) NewStatic() Static {
	return &static{}
}

func (s *static) Ping(key string) string {
	if key != "" {
		return fmt.Sprintf("$%d\r\n%s\r", len(key), key)
	}
	return domain.ResponsePong
}

func (s *static) Echo(arg string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
}
