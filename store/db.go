package store

import (
	"redis_app/domain"
)

// DB is the struct for data store
type DB struct {
	Single *single
	List   *list
}

type single struct {
	singleDB map[string]domain.Single
}
type list struct {
	listDB map[string]domain.ListValue
}

// NewDB returns the new data base for the session
func NewDB() *DB {
	return &DB{
		Single: &single{singleDB: make(map[string]domain.Single)},
		List:   &list{listDB: make(map[string]domain.ListValue)},
	}
}
