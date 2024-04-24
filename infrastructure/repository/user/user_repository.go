package repository

import (
	"context"

	userDomain "github.com/utya1414/blog-server/domain/user"
	"github.com/utya1414/blog-server/infrastructure/db"
)

type userRepository struct{
	store db.Store
}

func NewUserRepository(store db.Store) userDomain.UserRepository {
	return &userRepository{store: store}
}

func (r *userRepository) CreateUser(ctx context.Context, u *userDomain.CreateUser) error {
	_, err := r.store.CreateUser(ctx, db.CreateUserParams{
		Username:        u.GetUsername(),
		Email:           u.GetEmail(),
		HasshedPassword: u.GetPassword(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUser(ctx context.Context, username string) (*userDomain.User, error) {
	dbUser, err := r.store.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	
	domainUser, err := userDomain.NewUser(
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

// func (r *userRepository) ListUsers(ctx context.Context, limit int32, offset int32) ([]*user.User, error) {
// 	dbUsers, err := r.store.ListUsers(ctx, db.ListUsersParams{
// 		Limit: limit,
// 		Offset: offset,
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
