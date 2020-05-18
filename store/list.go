package store

import (
	"redis_app/domain"
	"sync"
)

type list struct {
	mutex  sync.Mutex
	listDB map[string]domain.ListValue
}

// GetList is a getList
func (db *list) GetList(key string) (*domain.ListValue, error) {
	return nil, nil
}

// SetList set the data from given args
func (db *list) SetList(args []string) error {
	return nil
}
