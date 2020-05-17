package parser

import (
	"fmt"
	"redis_app/command"
	"redis_app/domain"
)

func (h *singleCommandHandler) Handle(pr *parser) domain.RespString {
	key := pr.Arguments[0]
	fmt.Println(pr.Arguments)
	switch pr.Directive {
	case command.Get:
		return h.Get(key)
	case command.Set:
		if pr.Len < 3 {
			return domain.RespErrorWrongArgumentNumber
		}
		value := pr.Arguments[1]
		options := pr.Arguments[2:]
		return h.Set(key, value, options)
	case command.IncrBy:
		value := pr.Arguments[1]
		return h.IncrBy(key, value)
	default:
		return domain.RespErrorWrongSyntax
	}
}
