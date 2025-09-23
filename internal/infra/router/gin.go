package router

import (
	"go-monitor-tool/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

type (
	UserHandlers struct {
		Create   *handler.CreateUserHandler
		FindByID *handler.FindUserByIDHandler
		Search   *handler.SearchUserHandler
		Update   *handler.UpdateUserHandler
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
		usersApi := api.Group("/users")
		{
			usersApi.POST("", users.Create.Handle)
			usersApi.GET("/:id", users.FindByID.Handle)
			usersApi.GET("/search", users.Search.Handle)
			usersApi.GET("/:id/monitors", monitors.FindByUserID.Handle)
			usersApi.POST("/:id", users.Update.Handle)

		}

		monitorsApi := api.Group("/monitors")
		{
			monitorsApi.POST("", monitors.Create.Handle)
			monitorsApi.GET("/:id", monitors.FindByID.Handle)

		}

	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
