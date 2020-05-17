package parser

import (
	"redis_app/command"
	"redis_app/domain"
)

func (s *staticCommandHandler) Handle(pr *respRequest) domain.RespString {
	switch pr.Directive {
	case command.Ping:
		key := ""
		if len(pr.Arguments) != 0 {
			key = pr.Arguments[0]
		}
		return s.Ping(key)
	case command.Echo:
		arg := pr.Arguments[0]
		return s.Echo(arg)
	default:
		return domain.RespErrorWrongSyntax
	}
}
