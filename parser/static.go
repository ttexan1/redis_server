package parser

import (
	"redis_app/command"
	"redis_app/domain"
	"redis_app/usecase"
)

type staticCommandHandler struct {
	usecase.Static
}

func (s *staticCommandHandler) Handle(pr *respRequest) domain.RespString {
	switch pr.Directive {
	case command.Ping:
		if len(pr.Args) >= 2 {
			return domain.RespErrorWrongArgumentNumber
		}
		key := ""
		if len(pr.Args) == 1 {
			key = pr.Args[0]
		}
		return s.Ping(key)
	case command.Echo:
		if len(pr.Args) != 1 {
			return domain.RespErrorWrongArgumentNumber
		}
		arg := pr.Args[0]
		return s.Echo(arg)
	default:
		return domain.RespErrorWrongSyntax
	}
}
