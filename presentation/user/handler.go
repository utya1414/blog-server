package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utya1414/blog-server/application/user"
)

type handler struct {
	createUesrUseCase *user.CreateUserUseCase
}

type CreateUserParams struct {
    Username string `json:"username" binding:"required,alphanumunderscore,min=1,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func NewHandler(createUesrUseCase *user.CreateUserUseCase) *handler {
	return &handler{createUesrUseCase: createUesrUseCase}
}

func (h *handler) CreateUser(ctx *gin.Context) {
	var params CreateUserParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// todo: パスワードのハッシュ化
	err := h.createUesrUseCase.CreateUser(&user.CreateUserDto{
		Username: params.Username,
		Email: params.Email,
		Password: params.Password,
	})
	// todo: エラーハンドリング
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "ユーザーを作成しました"})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}