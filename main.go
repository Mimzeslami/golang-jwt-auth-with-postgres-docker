package main

import (
	"app/controllers"
	"app/database"
	"app/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DSN")
	database.Connect(dsn)
	database.Migrate()
	router := initRouter()
	router.Run(":8081")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1.0")
	{
		api.POST("/auth/login", controllers.LoginUser)
		api.POST("/auth/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
