package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ProxyMiddleware(router *gin.Engine) error {
	err := router.SetTrustedProxies([]string{"192.168.1.1", "192.168.1.2"})
	if err != nil {
		return err
	}

	return nil
}
