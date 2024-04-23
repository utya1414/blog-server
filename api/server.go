package api

import (
	"github.com/gin-gonic/gin"
	userApp "github.com/utya1414/blog-server/application/user"
	"github.com/utya1414/blog-server/infrastructure/db"
	"github.com/utya1414/blog-server/infrastructure/repository"
	userPre "github.com/utya1414/blog-server/presentation/user"
)


type Server struct {
	
}
// todo: ルーターの設定
func NewServer(store db.Store) (*gin.Engine, error) {
	server := setupRouter(store)
	return server, nil
}

func setupRouter(store db.Store) *gin.Engine {
	router := gin.Default()
	userRep := repository.NewUserRepository(store)
	userUse := userApp.NewCreateUserUseCase(userRep)
	userHandler := userPre.NewHandler(userUse)
	router.POST("/users", userHandler.CreateUser)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	return router
}
