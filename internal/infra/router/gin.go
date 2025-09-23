package router

import (
	"go-monitor-tool/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

type (
	UserHandlers struct {
		Create      *handler.CreateUserHandler
		FindByID    *handler.FindUserByIDHandler
		FindByEmail *handler.FindUserByEmailHandler
	}

	MonitorHandlers struct {
		Create       *handler.CreateMonitorHandler
		FindByID     *handler.FindMonitorByIDHandler
		FindByUserID *handler.FindMonitorsByUserHandler
	}
)

func NewGinRouter(users UserHandlers, monitors MonitorHandlers) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	{
		// Users
		api.POST("/users", users.Create.Handle)
		api.GET("/users/:id", users.FindByID.Handle)
		api.GET("/users/email/:email", users.FindByEmail.Handle)

		// Monitors
		api.POST("/monitors", monitors.Create.Handle)
		api.GET("/monitors/:id", monitors.FindByID.Handle)
		api.GET("/monitors/:user_id/monitors", monitors.FindByUserID.Handle)

	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
