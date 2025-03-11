package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(env string) gin.HandlerFunc {
	if env == "development" {
		return cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
		})
	}

	if env == "production" {
		return cors.New(cors.Config{
			AllowAllOrigins:  false,
			AllowCredentials: true,
			AllowOrigins:     []string{""},
		})
	}

	return nil
}
