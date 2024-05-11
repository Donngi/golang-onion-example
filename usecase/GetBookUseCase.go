package usecase

import "github.com/Donngi/golang-onion-example/domain"

type GetBookUseCase struct {
	repository domain.BookRepository
}

func NewGetBookUseCase(repository domain.BookRepository) *GetBookUseCase {
	return &GetBookUseCase{repository: repository}
}

func (uc *GetBookUseCase) Run(id string) (*domain.Book, error) {
	res, err := uc.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
