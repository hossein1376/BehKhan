package entities

type Book struct {
	ID    BookID
	Title string
}

type BookID int64

func (id BookID) ToInt64() int64 {
	return int64(id)
}
