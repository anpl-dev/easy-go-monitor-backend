package router

import (
	"go-monitor-tool/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(
	monitorHandler *handler.MonitorHandler,
	// userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()

	// Monitors
	r.POST("/monitors", monitorHandler.CreateMonitor)

	// Users
	// r.POST("/users", userHandler.CreateUser)

	// Health Check
	r.GET("/heahth", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	return r
}
