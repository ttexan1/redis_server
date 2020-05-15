package parser

import (
	"fmt"

	"redis_app/domain"
)

// Get fetch the value of given key
func (pr *parser) Set() string {
	if pr.Len <= 2 {
		return "INVALID REQUEST"
	}
	fmt.Println(pr.Arguments)
	err := pr.Store.Set(pr.Arguments)
	if err != nil {
		return domain.ErrorTypeNilValue
	}
	return domain.ResponseOK
}
