package router

import (
	"go-monitor-tool/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(
	userHandler *handler.UserHandler,
	monitorHandler *handler.MonitorHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	{
		// Users
		api.POST("/users", userHandler.CreateUser)

		// Monitors
		api.POST("/monitors", monitorHandler.CreateMonitor)

	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
