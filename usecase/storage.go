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

type Single struct {
}

type List struct {
}

type UserCase struct {
	Storage
}

func NewUseCase(s Storage) *UserCase {
	return &UserCase{s}
}
