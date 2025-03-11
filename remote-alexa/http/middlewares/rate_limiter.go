package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5)

func RateLimiter(ctx *gin.Context) {
	if !limiter.Allow() {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
