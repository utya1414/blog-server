package userRep

import (
	"time"

	userDomain "github.com/utya1414/blog-server/domain/user"
)

type UserDto struct {
	User *userDomain.User
	CreatedAt time.Time
	UpdatedAt time.Time
}