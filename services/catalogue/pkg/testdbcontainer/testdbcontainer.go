package testdbcontainer

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

func New(ctx context.Context) (*mariadb.MariaDBContainer, error) {
	return mariadb.RunContainer(ctx,
		testcontainers.WithImage("mariadb:11.3.2"),
	)
}
