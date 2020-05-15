package store

import (
	"fmt"

	"redis_app/command"
	"redis_app/domain"
)

// DB is the struct for data store
type DB struct {
	St map[string]*domain.Single
}

// // DB is the key-value store for the current session
// type DB struct {
// 	Single map[string]*domain.Single
// 	List   map[string]*domain.ListValue
// }

// // NewDB returns the new data base for the session
// func NewDB() *DB {
// 	return &DB{
// 		Single: make(map[string]*domain.Single),
// 		List:   make(map[string]*domain.ListValue),
// 	}
// }

// Get returns the value of given key
func (db *DB) Get(key string) (*domain.Single, error) {
	info, ok := db.St[key]
	if !ok {
		return &domain.Single{}, fmt.Errorf("record not found")
	}
	return info, nil
}

// Set set the data from given args
func (db *DB) Set(args []string) error {
	key := args[0]
	info, ok := db.St[key]

	options := args[2:]
	// fmt.Println(options)
	for _, option := range options {
		switch option {
		case command.SetOptionNX:
			if ok {
				return fmt.Errorf("Invalid")
			}
			info = &domain.Single{
				Value:  args[1],
				Length: len(args[1]),
			}
			db.St[key] = info
			return nil
		case command.SetOptionXX:
			if !ok {
				return fmt.Errorf("Invalid")
			}
			info.Value = args[1]
			info.Length = len(args[1])
			return nil
		}
	}
	info = &domain.Single{
		Value:  args[1],
		Length: len(args[1]),
	}
	db.St[key] = info
	return nil
}

func timeSecond(millisecond int) int {
	return millisecond * 1000
}
