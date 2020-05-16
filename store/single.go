package store

import (
	"fmt"
	"math/rand"
	"redis_app/domain"
)

// GetValue is a getvalue
func (ss *single) GetValue(key string) (*domain.Single, error) {
	rd := []string{"AA", "BB", "CC", "DD", "EE"}[rand.Intn(5)]
	fmt.Println(rd)
	ss.singleDB[rd] = &domain.Single{
		Value: "aaaaa",
	}
	fmt.Println(ss.singleDB)
	return &domain.Single{
		Value: "aaaaa",
	}, nil
}

// SetValue set the data from given args
func (ss *single) SetValue(args []string) error {
	return nil
}

// Delete delte value
func (ss *single) Delete(string) {

}
