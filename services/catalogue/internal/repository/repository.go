package repository

type BookRepository interface {
	GetByID(...int64) ([]Book, error)
	GetAll() ([]Book, error)
}
