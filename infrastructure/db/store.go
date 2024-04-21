package db

import (
	"database/sql"

	_ "github.com/blog-server/infrastructure/db/postgresql/sqlc"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
	}
}