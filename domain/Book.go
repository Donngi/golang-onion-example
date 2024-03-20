package domain

import (
	"github.com/google/uuid"
)

type Book struct {
	id     string
	name   string
	author string
}

func NewBook(name string, author string) (*Book, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Book{
		id:     id.String(),
		name:   name,
		author: author,
	}, nil
}
