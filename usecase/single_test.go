package usecase

import (
	"redis_app/domain"
	"testing"
)

// type SingleStore interface {
// 	GetValue(string) (domain.Single, *domain.Error)
// 	SetValue(string, domain.Single) *domain.Error
// 	Delete(string)
// }

type testSingleStore struct{}

func (t *testSingleStore) GetValue(key string) (domain.Single, *domain.Error) {
	return domain.Single{Value: "string"}, nil
}

func (t *testSingleStore) SetValue(key string, value domain.Single) *domain.Error {
	return nil
}

func (t *testSingleStore) Delete(key string) {
	return
}
func (t *testSingleStore) IncrSet(string, int) (int, *domain.Error) {
	return 0, nil
}

func TestUsecase(t *testing.T) {
	s := &single{SingleStore: &testSingleStore{}}
	if got := s.Get("name"); got != domain.RespBulkString("string") {
		t.Fatalf("Error Expected: %v, Got: %v", "string", got)
	}
}
