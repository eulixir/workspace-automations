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
		log.Fatalf("Failed to load config: %v", err)
	}

	router := setupRouter()

	logger.Info("Starting server on port", zap.String("port", cfg.Port))
	router.Run(fmt.Sprintf(":%s", cfg.Port))
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.RateLimiter)

	routes.SetupRoutes(router)

	return router
}
