package repository

import (
	"context"

	"github.com/utya1414/blog-server/domain/user"
	"github.com/utya1414/blog-server/infrastructure/db"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {
	dbUser, err := db.Store.CreateUser(ctx, db.CreateUserParams{
		Username:        u.GetUsername(),
		Email:           u.GetEmail(),
		HasshedPassword: u.GetHasshedPassword(),
	})
	if err != nil {
		return nil, err
	}

	// TODO: dbUserがどういうフィールドを持っているか知っている実装なのでカプセル化したい。
	domainUser, err := user.NewUser(
		dbUser.Username, 
		dbUser.Email, 
		dbUser.HasshedPassword,
		dbUser.UpdatedAt.String(),
		dbUser.CreatedAt.String(),
	)
	if err != nil {
		return nil, err
	}

	return domainUser, nil
}

// func (r *userRepository) GetUser(ctx context.Context, username string) (*user.User, error) {
// 	dbUser, err := db.Store.GetUser(ctx, username)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	domainUser, err := user.NewUser(
// 		dbUser.Username, 
// 		dbUser.Email, 
// 		dbUser.HasshedPassword,
// 		dbUser.UpdatedAt.String(),
// 		dbUser.CreatedAt.String(),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return domainUser, nil
// }

// func (r *userRepository) ListUsers(ctx context.Context, limit int32) ([]*user.User, error) {
// 	dbUsers, err := db.Store.ListUsers(ctx, db.ListUsersParams{
// 		Limit: limit,
// 		Offset: 0,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var domainUsers []*user.User
// 	for _, dbUser := range dbUsers {
// 		domainUser, err := user.NewUser(
// 			dbUser.Username, 
// 			dbUser.Email, 
// 			dbUser.HasshedPassword,
// 			dbUser.UpdatedAt.String(),
// 			dbUser.CreatedAt.String(),
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		domainUsers = append(domainUsers, domainUser)
// 	}

// 	return domainUsers, nil
// }
