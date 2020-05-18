package usecase

// Usecase is the struct for usecase
type Usecase struct {
	SingleStore SingleStore
	ListStore   ListStore
}

// NewUsecase returns initialized struct with database pointer
func NewUsecase(s SingleStore, l ListStore) *Usecase {
	return &Usecase{
		SingleStore: s,
		ListStore:   l,
	}
}
