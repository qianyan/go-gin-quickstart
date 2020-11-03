package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	"github.com/qianyan/go-gin-quickstart/domain/users"
	"github.com/qianyan/go-gin-quickstart/infra"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate(db)
}

func main() {

	db := infra.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

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
