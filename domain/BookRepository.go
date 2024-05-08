package domain

type BookRepository interface {
	Add(book *Book) (*Book, error)
	Get(id string) (*Book, error)
}
