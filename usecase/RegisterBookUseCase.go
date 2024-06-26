package usecase

import (
	"log"

	"github.com/Donngi/golang-onion-example/domain"
)

type RegisterBookUseCase struct {
	repository domain.BookRepository
}

func NewRegisterBookUseCase(repository domain.BookRepository) *RegisterBookUseCase {
	return &RegisterBookUseCase{repository: repository}
}

func (uc *RegisterBookUseCase) Run(name string, author string) (*domain.Book, error) {
	book, err := domain.NewBook(name, author)
	if err != nil {
		log.Fatalf("Failed to create book: %v", err)
	}

	res, err := uc.repository.Add(book)
	if err != nil {
		log.Fatalf("Failed to add book: %v", err)
	}

	return res, nil
}
