package parser

import (
	"redis_app/command"
	"redis_app/domain"
	"redis_app/usecase"
)

type singleCommandHandler struct {
	pr *respRequest
	usecase.Single
}

func (h *singleCommandHandler) Handle(pr *respRequest) domain.RespString {
	if len(pr.Args) == 0 {
		return domain.RespErrorWrongArgumentNumber
	}
	key := pr.Args[0]
	switch pr.Directive {
	case command.Get:
		return h.Get(key)
	case command.Set:
		if len(pr.Args) < 2 {
			return domain.RespErrorWrongArgumentNumber
		}
		value := pr.Args[1]
		var options []string
		if len(pr.Args) >= 2 {
			options = pr.Args[2:]
		}
		return h.Set(key, value, options)
	case command.IncrBy:
		if len(pr.Args) != 2 {
			return domain.RespErrorWrongArgumentNumber
		}
		value := pr.Args[1]
		return h.IncrBy(key, value)
	case command.DecrBy:
		if len(pr.Args) != 2 {
			return domain.RespErrorWrongArgumentNumber
		}
		value := pr.Args[1]
		return h.DecrBy(key, value)
	case command.Incr:
		return h.Incr(key)
	case command.Decr:
		return h.Decr(key)
	case command.Del:
		return h.Del(pr.Args)
	default:
		return domain.RespErrorWrongSyntax
	}
}
