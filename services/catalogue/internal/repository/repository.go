package repository

type BookRepository interface {
	GetByID(int) (*Book, error)
	GetAll() ([]Book, error)
}
