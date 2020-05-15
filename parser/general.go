package parser

import (
	"fmt"
	"strconv"
	"strings"

	"redis_app/command"
	"redis_app/domain"
	"redis_app/store"
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

// ParseCommand parse the given text to response string
func ParseCommand(text string, db *store.DB) string {
	pr := rawStringToArguments(text)
	fmt.Println(pr)
	pr.Store = db

	switch pr.Directive {
	case command.Ping:
		return pr.Ping()
	case command.Get:
		return pr.Get()
	case command.Set:
		return pr.Set()
	default:
		return "Invalid Command"
	}
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
			Directive: command.InvalidFormat,
		}
	}
	l, err := strconv.Atoi(strings.Split(args[0], "*")[1])
	if err != nil {
		return &parser{
			Directive: command.InvalidFormat,
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
