package repository

import (
	"context"

	domainUser "github.com/utya1414/blog-server/domain/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domainUser.CreateUser) error
	GetUser(ctx context.Context, username string) (*domainUser.User, error)
	// ListUsers(ctx context.Context, limit int32, offset int32) ([]*domainUser.User, error)
}