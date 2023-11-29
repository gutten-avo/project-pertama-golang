package routes

import (
	"net/http"
	user_controller "project-pertama-golang/controller/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS", "TRACE", "CONNECT"},
		AllowHeaders:     []string{"Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "Content-Type", "Content-Length", "Date", "origin", "Origins", "x-requested-with", "access-control-allow-methods", "access-control-allow-credentials", "apikey"},
		ExposeHeaders:    []string{"Content-Length"},
	}))

	r.GET("/ping", ping)
	r.GET("/get-users", user_controller.GetAllUser)
	r.POST("/add-user", user_controller.AddUser)
	return r
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
