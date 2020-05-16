package store

import (
	"redis_app/domain"
	"testing"
)

func TestSingleSet(t *testing.T) {
	ss := &single{singleDB: map[string]domain.Single{}}
	keyValues := map[string]domain.Single{
		"key1": {Value: "value1", Length: 6},
		"key2": {Value: "value2", Length: 6},
		"key3": {Value: "value3", Length: 6},
		"key4": {Value: "value4", Length: 6},
	}
	for key, value := range keyValues {
		ss.SetValue(key, value)
	}
	for key, value := range keyValues {
		t.Run(key, func(t *testing.T) {
			if got, err := ss.GetValue(key); err != nil || got.Value != value.Value {
				t.Errorf("Wrong")
			}
		})
	}
}

func TestSingleDelete(t *testing.T) {
	key := "key1"
	ss := &single{singleDB: map[string]domain.Single{
		key: {Value: "Drink", Length: 5},
	}}
	ss.Delete(key)
	if _, ok := ss.singleDB[key]; ok {
		t.Fatalf("%s doesn't deleted", key)
	}
}

// *1
// $4
// ping/

// *2
// $3
// get
// $3
// key/

// *3
// $3
// set
// $3
// key
// $5
// value/
