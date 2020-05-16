package parser

import (
	"redis_app/command"
	"redis_app/domain"
)

// type singleCommandHandler struct {
// 	pr parser
// 	usecase.SingleStore
// }

func (h *singleCommandHandler) Handle() string {
	// h.Set()

	switch h.pr.Directive {
	case command.Ping:
		return h.pr.Ping()
	case command.Get:
		return h.Get(h.pr.Arguments[0])
	// case command.Set:
	// 	return h.pr.Set()
	default:
		return domain.ErrorTypeWrongSyntax
	}
}

// Get fetch the value of given key
// func (h *singleCommandHandler) Get() string {
// 	if h.pr.Len != 2 {
// 		return "INVALID REQUEST"
// 	}
// 	info, err := h.Get(h.pr.Arguments[0])
// 	if err != nil {
// 		return domain.ErrorTypeNilValue
// 	}
// 	return fmt.Sprintf("$%d\n%s\r\n", info.Length, info.Value)
// }
