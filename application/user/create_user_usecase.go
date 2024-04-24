package user

import (
	"context"

	userDomain "github.com/utya1414/blog-server/domain/user"
	errDomain "github.com/utya1414/blog-server/domain/user/error"
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
	if _, err := uc.userRep.GetUser(context.Background(), dto.Username); err == nil{
		return errDomain.ErrUserAlreadyExists
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
