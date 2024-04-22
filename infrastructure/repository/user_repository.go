package repository

import (
	"context"

	"github.com/utya1414/blog-server/domain/user"
	"github.com/utya1414/blog-server/infrastructure/db"
)

type userRepository struct{
	store db.Store
}

func NewUserRepository(store db.Store) UserRepository {
	return &userRepository{store: store}
}

func (r *userRepository) CreateUser(ctx context.Context, u *user.User) error {
	dbUser, err := r.store.CreateUser(ctx, db.CreateUserParams{
		Username:        u.GetUsername(),
		Email:           u.GetEmail(),
		HasshedPassword: u.GetHasshedPassword(),
	})
	if err != nil {
		return err
	}

	// TODO: dbUserがどういうフィールドを持っているか知っている実装なのでカプセル化したい。
	_, err = user.NewUser(
		dbUser.Username, 
		dbUser.Email, 
		dbUser.HasshedPassword,
		dbUser.UpdatedAt.String(),
		dbUser.CreatedAt.String(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUser(ctx context.Context, username string) (*user.User, error) {
	dbUser, err := r.store.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	
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
