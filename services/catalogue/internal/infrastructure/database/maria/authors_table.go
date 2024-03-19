package maria

import (
	"database/sql"
)

type AuthorsTable struct {
	tx *sql.Tx
}

func newAuthorsTable(tx *sql.Tx) AuthorsTable {
	return AuthorsTable{tx: tx}
}
