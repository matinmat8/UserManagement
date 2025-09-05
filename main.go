// @title Authentication API
// @version 1.0
// @description This is a sample authentication service with OTP + JWT in Go + Gin.
// @host localhost:8000
// @BasePath /api/v1

package main

import (
	"authentication/bootstrap"
	_ "authentication/docs"
	"authentication/middleware"
	"authentication/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.Use(middleware.ErrorHandling())
	app := bootstrap.InitAppContainer()
	routes.Urls(r, app)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(gin.DebugMode)
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
