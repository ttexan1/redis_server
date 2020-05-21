package store_test

import (
	"redis_app/domain"
	"redis_app/store"
	"testing"
)

func TestSingleSet(t *testing.T) {
	db := store.NewDB()
	ss := db.Single
	keyValues := map[string]domain.Single{
		"key1": {Value: "value1"},
		"key2": {Value: "value2"},
		"key3": {Value: "value3"},
		"key4": {Value: "value4"},
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
	db := store.NewDB()
	ss := db.Single
	ss.SetValue(key, domain.Single{Value: "value"})
	ss.Delete(key)
	if _, err := ss.GetValue(key); err.Message != domain.RespErrorNilValue {
		t.Fatalf("%s doesn't deleted", key)
	}
}
