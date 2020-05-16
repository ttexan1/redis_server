package parser

import (
	"fmt"
	"strconv"
	"strings"

	"redis_app/command"
	"redis_app/domain"
	"redis_app/store"
	"redis_app/usecase"
)

// Store is the interface
type Store interface {
	Get(string) (*domain.Single, error)
	Set([]string) error
}

type parser struct {
	Directive string
	Arguments []string
	// Len is a length for the top command
	Len int

	Store Store
}
type singleCommandHandler struct {
	pr *parser
	usecase.Single
}

type listCommandHandler struct {
	pr *parser
	usecase.List
}

type ParseHandler interface {
	Handle() string
}

// InitParser registers the parser
func InitParser(uc *usecase.UseCase) map[string]ParseHandler {
	handlers := map[string]ParseHandler{
		"single": &singleCommandHandler{
			Single: uc.NewSingle(),
		},
	}
	return handlers
	// handler := listCommandHandler{
	// 	List: uc.NewList(),
	// }
}

// ParseCommand parse the given text to response string
func ParseCommand(text string, db *store.DB, uc *usecase.UseCase) string {
	parsers := map[string]ParseHandler{}
	pr := rawStringToArguments(text)
	pr.Store = db

	fmt.Println(pr)
	if _, ok := command.StaticCommandList[pr.Directive]; ok {
	}
	if _, ok := command.SingleCommandWhiteList[pr.Directive]; ok {
		parsers["single"].Handle()
		handler := singleCommandHandler{
			pr:     pr,
			Single: uc.NewSingle(),
		}
		result := handler.Handle()
		fmt.Println(result)
		return result
	}
	if _, ok := command.ListCommandWhiteList[pr.Directive]; ok {

	}
	return domain.ErrorTypeUnknownCommand
}

// example
// `*2
// $4
// ping
// $3
// aaa`

func rawStringToArguments(text string) *parser {
	pr := &parser{}
	args := strings.Split(text, "\n")
	fmt.Println(args)
	if len(args) < 3 {
		return &parser{
			Directive: "Invalid",
		}
	}
	l, err := strconv.Atoi(strings.Split(args[0], "*")[1])
	if err != nil {
		return &parser{
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
