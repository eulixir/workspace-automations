package main

import (
	"fmt"
	"log"

	"github.com/eulixir/workspace-automations/cmd/alexa-webhook/config"
	"github.com/eulixir/workspace-automations/http/middlewares"
	"github.com/eulixir/workspace-automations/http/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config", zap.Error(err))
	}

	router, err := setupRouter(cfg)
	if err != nil {
		logger.Error("Failed to setup router", zap.Error(err))
	}

	logger.Info("Starting server on port", zap.String("port", cfg.Port))
	router.Run(fmt.Sprintf(":%s", cfg.Port))
}

func setupRouter(cfg *config.Config) (*gin.Engine, error) {
	router := gin.Default()

	err := middlewares.ProxyMiddleware(router)
	if err != nil {
		return nil, err
	}

	router.Use(middlewares.RateLimiter)
	router.Use(middlewares.CORSMiddleware(cfg.Env))

	routes.SetupRoutes(router)

	return router, nil
}
