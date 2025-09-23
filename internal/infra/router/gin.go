package router

import (
	"go-monitor-tool/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	Create      *handler.CreateUserHandler
	FindByID    *handler.FindUserByIDHandler
	FindByEmail *handler.FindUserByEmailHandler
}

type MonitorHandlers struct {
	Create *handler.CreateMonitorHandler
}

func NewGinRouter(users UserHandlers, monitors MonitorHandlers) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	{
		// Users
		api.GET("/users/:id", users.FindByID.Handle)
		api.GET("/users/:email", users.FindByEmail.Handle)
		api.POST("/users", users.Create.Handle)

		// Monitors
		api.POST("/monitors", monitors.Create.Handle)

	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
