// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
}

var _ Querier = (*Queries)(nil)
