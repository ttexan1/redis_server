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
	Args      []string
	Len       int
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
func InitParser(uc *usecase.Usecase) map[string]ParseHandler {
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

// HandleRequest parse the given text to response string
func HandleRequest(text string, parsers map[string]ParseHandler) domain.RespString {
	pr, err := decode(text)
	if err != nil {
		return domain.RespErrorWrongSyntax
	}
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

func decode(text string) (*respRequest, error) {
	pr := &respRequest{}
	texts := strings.Split(text, "\r\n")
	texts = texts[:len(texts)-1]
	if len(texts) < 3 {
		return nil, fmt.Errorf("Invalid Syntax")
	}
	reqHeader := texts[0]
	if len(reqHeader) < 1 {
		return nil, fmt.Errorf("Invalid Syntax")
	}
	// check the request length
	l, err := strconv.Atoi(reqHeader[1:])
	if err != nil {
		return nil, fmt.Errorf("Invalid Syntax")
	}
	pr.Len = l

	pr.Directive = toLowerCase(texts[2])
	if pr.Len >= 2 {
		isInfo := true
		bsLength := 0
		for _, arg := range texts[3:] {
			if isInfo {
				// if strings.HasPrefix(arg, "*") {}
				if strings.HasPrefix(arg, "$") {
					i, err := strconv.Atoi(arg[1:])
					if err != nil {
						return nil, fmt.Errorf("Invalid Syntax")
					}
					bsLength = i
					isInfo = false
					continue
				}
			}
			pr.Args = append(pr.Args, arg[:bsLength])
			isInfo = true
		}
	}
	return pr, nil
}

func toLowerCase(text string) string {
	b := []byte(text)
	for i, c := range []byte(text) {
		if 'A' <= c && c <= 'Z' {
			b[i] += 'a' - 'A'
		}
	}
	return string(b)
}
