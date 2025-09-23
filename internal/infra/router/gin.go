package router

import (
	"go-monitor-tool/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(
	monitorHandler *handler.MonitorHandler,
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	{
		// Monitors
		api.POST("/monitors", monitorHandler.CreateMonitor)

		// Users
		api.POST("/users", userHandler.CreateUser)

	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
