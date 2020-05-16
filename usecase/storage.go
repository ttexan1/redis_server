package usecase

// UseCase is the struct for usecase
type UseCase struct {
	SingleStore SingleStore
	ListStore   ListStore
}

// NewUseCase returns initialized struct with database pointer
func NewUseCase(s SingleStore, l ListStore) *UseCase {
	return &UseCase{
		SingleStore: s,
		ListStore:   l,
	}
}
