package main

import (
	"database/sql"
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/utya1414/blog-server/api"
	"github.com/utya1414/blog-server/infrastructure/db"
	"github.com/utya1414/blog-server/util"
)

func validaterinit() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("alphanumunderscore", func(fl validator.FieldLevel) bool {
			return regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(fl.Field().String())
		})
	}
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	conn, err := sql.Open(config.DBDriver, config.DBSource);
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	router, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	validaterinit()
	
	err = router.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
