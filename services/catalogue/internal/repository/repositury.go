package repository

type BookRepository interface {
	GetAll() []Book
}
