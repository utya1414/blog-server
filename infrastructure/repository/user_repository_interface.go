package repository

import (
	"context"

	"github.com/utya1414/blog-server/domain/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *user.User) error
	// GetUser(ctx context.Context, username string) (*User, error)
	// ListUsers(ctx context.Context, limit int32) ([]*User, error)
}