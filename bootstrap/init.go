package bootstrap

import (
	v1 "authentication/controllers"
	"authentication/db"
	"authentication/repositories"
	"authentication/services"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type AppContainer struct {
	Redis   *redis.Client
	Limiter *redis_rate.Limiter
	AuthAPI v1.AuthAPI
}

func InitAppContainer() *AppContainer {
	redisClient := db.RedisClient()

	//jwtAuth := jwt.Jwt{}

	authRepo := repositories.NewAuthRepository(redisClient)
	authService := services.NewAuthService(authRepo)
	authController := v1.NewAuthAPI(authService)

	limiter := redis_rate.NewLimiter(redisClient)

	return &AppContainer{
		Redis:   redisClient,
		Limiter: limiter,
		AuthAPI: authController,
	}

}
