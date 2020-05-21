package store

import (
	"redis_app/domain"
	"sync"
)

// DB is the struct for data store
type DB struct {
	Single *single
	List   *list
	// KeyType *keyType
}

// type keyType struct {
// 	keyDB map[string]string
// }

// NewDB returns the new data base for the session
func NewDB() *DB {
	return &DB{
		// KeyType: &keyType{keyDB: make(map[string]string)},
		Single: &single{singleDB: make(map[string]domain.Single), mutex: sync.Mutex{}},
		List:   &list{listDB: make(map[string]domain.ListValue)},
	}
}
