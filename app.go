package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qianyan/go-gin-quickstart/infra/config"
	"github.com/qianyan/go-gin-quickstart/infra/logger"

	"github.com/qianyan/go-gin-quickstart/domain/users"
	"github.com/qianyan/go-gin-quickstart/infra"
)

func main() {

	db := &infra.Sqlite{}
	db.OpenDB("./gorm.db")
	users.Init(db, true)
	defer db.CloseDB()

	r := gin.Default()

	config.LoadConfig("config.json")

	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	gin.SetMode(config.Conf.Mode)

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	userResource(r)

	healthCheck(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func userResource(r *gin.Engine) {
	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(true))
	users.SpecifiedUser(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
}

func healthCheck(r *gin.Engine) {
	healthy := r.Group("/api/ping")
	healthy.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
