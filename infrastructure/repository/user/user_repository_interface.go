package userRep

import (
	"context"

	userDomain "github.com/utya1414/blog-server/domain/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *userDomain.User) error
	GetUser(ctx context.Context, username string) (*UserDto, error)
	// ListUsers(ctx context.Context, limit int32, offset int32) ([]*domainUser.User, error)
}