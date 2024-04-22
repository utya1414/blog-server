package user

import (
	"context"

	"github.com/pkg/errors"
	userDomain "github.com/utya1414/blog-server/domain/user"
	"github.com/utya1414/blog-server/infrastructure/repository"
)

type CreateUserUseCase struct {
	userRep repository.UserRepository
}

func NewCreateUserUseCase(userRep repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRep: userRep}
}

type CreateUserDto struct {
	Username string
	Email string
	Password string
}

func (uc *CreateUserUseCase) CreateUser(dto *CreateUserDto) error {
	if _, err := uc.userRep.GetUser(context.Background(), dto.Username); err != nil{
		return errors.New("ユーザーが既に存在しています")
	}
	
	// ドメインモデルを生成
	u, err := userDomain.NewCreateUser(dto.Username, dto.Email, dto.Password)
	if err != nil {
		return err
	}

	err = uc.userRep.CreateUser(context.Background(), u)
	if err != nil {
		return err
	}

	return nil
}
