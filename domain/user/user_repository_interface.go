package user

import "github.com/utya1414/blog-server/infrastructure/db"

type UserRepository interface {
	CreateUser(user *db.CreateUserParams) (*User, error) 
	GetUser(username string) (*User, error)
	ListUsers(limit int32) ([]*User, error)
}