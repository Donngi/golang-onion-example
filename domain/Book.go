package domain

type Book struct {
	id     string
	name   string
	author string
}

func NewBook(id string, name string, author string) *Book {
	return &Book{
		id:     id,
		name:   name,
		author: author,
	}
}
