package usecase

import "redis_app/domain"

// Storage is the interface for the store
type Storage interface {
	SingleStore
	ListStore
}

// SingleStore is the interface
type SingleStore interface {
	GetValue(string) (*domain.Single, error)
	SetValue([]string) error
	Delete(string)
}

type ListStore interface {
	GetList(string) (*domain.ListValue, error)
	SetList([]string) error
}

//

//
//
//
//
//
//

type Single interface {
	Get(string) string
	Set(string, []string) string
	Del(string) string
}
type single struct {
	SingleStore
}

func (uc *UseCase) NewSingle() Single {
	return &single{
		SingleStore: uc.SingleStore,
	}
}

func (sg *single) Get(key string) string {
	if _, err := sg.GetValue(key); err != nil {

	}
	return "+PONG"
}
func (sg *single) Set(key string, options []string) string { return "" }
func (sg *single) Del(key string) string                   { return "" }

//
//
//
//

type List struct {
}

type useCase struct {
}

type UseCase struct {
	SingleStore SingleStore
	ListStore   ListStore
}

// NewUseCase returns initialized struct
// store.DB is given on Storage
func NewUseCase(s SingleStore, l ListStore) *UseCase {
	return &UseCase{
		SingleStore: s,
		ListStore:   l,
	}
}
