package db

import (
	"database/sql"

	sqlc "github.com/utya1414/blog-server/infrastructure/db/postgresql/sqlc"
)

type Store interface {
	sqlc.Querier
}

type SQLStore struct {
	*sqlc.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
		Queries: sqlc.New(db),
	}
}

// sqlcパッケージの隠蔽
type CreateUserParams = sqlc.CreateUserParams
type ListUserParams = sqlc.ListUsersParams