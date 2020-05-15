package parser

import "redis_app/domain"

func (pr *parser) Ping() string {
	if pr.Len == 1 {
		return domain.ResponsePong
	}
	if pr.Len == 2 {
		return pr.Arguments[0]
	}
	return "Invalid Length Of Argument"
}
