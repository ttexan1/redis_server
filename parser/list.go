package parser

import "redis_app/usecase"

type listCommandHandler struct {
	pr *respRequest
	usecase.List
}
