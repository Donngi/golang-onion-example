package domain

import (
	"github.com/google/uuid"
)

type Book struct {
	Id     string
	Name   string
	Author string
}

func NewBook(name string, author string) (*Book, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Book{
		Id:     id.String(),
		Name:   name,
		Author: author,
	}, nil
}
