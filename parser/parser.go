package parser

import (
	"fmt"
	"strconv"
	"strings"

	"redis_app/command"
	"redis_app/domain"
	"redis_app/usecase"
)

type respRequest struct {
	Directive string
	Arguments []string
	// Len is a length for the top command
	Len int
}

type staticCommandHandler struct {
	usecase.Static
}
type singleCommandHandler struct {
	pr *respRequest
	usecase.Single
}

type listCommandHandler struct {
	pr *respRequest
	usecase.List
}

// ParseHandler is the interface for parser
type ParseHandler interface {
	Handle(*respRequest) domain.RespString
}

const (
	single = "single"
	static = "static"
	list   = "list"
)

// InitParser registers the parser
func InitParser(uc *usecase.UseCase) map[string]ParseHandler {
	handlers := map[string]ParseHandler{
		single: &singleCommandHandler{
			Single: uc.NewSingle(),
		},
		static: &staticCommandHandler{
			Static: uc.NewStatic(),
		},
	}
	return handlers
}

// ParseCommand parse the given text to response string
func ParseCommand(text string, parsers map[string]ParseHandler) domain.RespString {
	pr := rawStringToArguments(text)
	if _, ok := command.StaticCommandList[pr.Directive]; ok {
		return parsers[static].Handle(pr)
	}
	if _, ok := command.SingleCommandList[pr.Directive]; ok {
		return parsers[single].Handle(pr)
	}
	if _, ok := command.ListCommandWhiteList[pr.Directive]; ok {
	}
	return domain.RespErrorUnknownCommand
}

// example
// `*2
// $4
// ping
// $3
// aaa`

func rawStringToArguments(text string) *respRequest {
	pr := &respRequest{}
	args := strings.Split(text, "\r\n")
	fmt.Println(args)
	if len(args) < 3 {
		return &respRequest{
			Directive: "Invalid",
		}
	}
	l, err := strconv.Atoi(strings.Split(args[0], "*")[1])
	if err != nil {
		return &respRequest{
			Directive: "Invalid",
		}
	}
	pr.Len = l

	pr.Directive = low(args[2])
	if pr.Len >= 2 {
		for _, arg := range args[3:] {
			if strings.HasPrefix(arg, "$") {
				continue
			}
			pr.Arguments = append(pr.Arguments, arg)
		}
	}
	return pr
}

func low(text string) string {
	return strings.ToLower(text)
}
