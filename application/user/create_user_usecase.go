package user

import (
	"context"

	errDomain "github.com/utya1414/blog-server/domain/error"
	userDomain "github.com/utya1414/blog-server/domain/user"
	userRep "github.com/utya1414/blog-server/infrastructure/repository/user"
)

type CreateUserUseCase struct {
	userRep userRep.UserRepository
}

func NewCreateUserUseCase(userRep userRep.UserRepository) *CreateUserUseCase {
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
	u, err := userDomain.NewUser(dto.Username, dto.Email, dto.Password)
	if err != nil {
		return err
	}

	err = uc.userRep.CreateUser(context.Background(), u)
	if err != nil {
		return err
	}

	return nil
}
