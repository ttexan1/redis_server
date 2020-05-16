package usecase

import "redis_app/domain"

// List is the interface for list command
type List interface {
}

// ListStore is the interface
type ListStore interface {
	GetList(string) (*domain.ListValue, error)
	SetList([]string) error
}
