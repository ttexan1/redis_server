package parser

import (
	"fmt"
	"redis_app/command"
	"redis_app/domain"
)

func (h *singleCommandHandler) Handle(pr *parser) string {
	key := pr.Arguments[0]
	fmt.Println(pr.Arguments)
	switch pr.Directive {
	case command.Get:
		return h.Get(key)
	case command.Set:
		value := pr.Arguments[1]
		options := pr.Arguments[2:]
		return h.Set(key, value, options)
	case command.IncrBy:
		value := pr.Arguments[1]
		return h.IncrBy(key, value)
	default:
		return domain.ErrorTypeWrongSyntax
	}
}
