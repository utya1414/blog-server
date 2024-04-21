package user

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error) 
	// GetUser(ctx context.Context, username string) (*User, error)
	// ListUsers(ctx context.Context, limit int32) ([]*User, error)
}