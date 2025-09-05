package routes

import (
	"authentication/bootstrap"
	"github.com/gin-gonic/gin"
)

func Urls(r *gin.Engine, app *bootstrap.AppContainer) *gin.Engine {
	apiV1 := r.Group("api/v1/auth/")
	{
		auth := apiV1.Group("")
		{
			auth.POST("/login/", app.AuthAPI.Login)
			auth.POST("/send/otp/", app.AuthAPI.SendOTP)
			auth.GET("/profile/", app.AuthAPI.Profile)
			auth.GET("/users", app.AuthAPI.ListUsers)
		}
	}

	// example of protected routes with jwt token

	//protected := r.Group("api/v1/user/")
	//protected.Use(middleware.JWTAuthMiddleware())
	//{
	//	protected.GET("/profile/", app.AuthAPI.Profile)
	//}

	return r
}
