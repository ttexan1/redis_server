package parser

import (
	"fmt"

	"redis_app/domain"
)

// Get fetch the value of given key
func (pr *parser) Get() string {
	if pr.Len != 2 {
		return "INVALID REQUEST"
	}
	info, err := pr.Store.Get(pr.Arguments[0])
	if err != nil {
		return domain.ErrorTypeNilValue
	}
	return fmt.Sprintf("$%d\n%s\r\n", info.Length, info.Value)
}
