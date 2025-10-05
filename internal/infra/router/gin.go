package router

import (
	uhandler "go-monitor-tool/internal/user/adapter/handler"
	mhandler "go-monitor-tool/internal/monitor/adapter/handler"

	"github.com/gin-gonic/gin"
)

type (
	UserHandlers struct {
		Create   *uhandler.CreateUserHandler
		FindByID *uhandler.FindUserByIDHandler
		Search   *uhandler.SearchUsersHandler
		Update   *uhandler.UpdateUserHandler
		Delete   *uhandler.DeleteUserHandler
	}

	MonitorHandlers struct {
		Create   *mhandler.CreateMonitorHandler
		FindByID *mhandler.FindMonitorByIDHandler
		Search   *mhandler.SearchMonitorsHandler
		Update   *mhandler.UpdateMonitorHandler
		Delete   *mhandler.DeleteMonitorHandler
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
			usersApi.PUT("/:id", users.Update.Handle)
			usersApi.DELETE("/:id", users.Delete.Handle)
		}

		monitorsApi := api.Group("/monitors")
		{
			monitorsApi.POST("", monitors.Create.Handle)
			monitorsApi.GET("/:id", monitors.FindByID.Handle)
			usersApi.GET("/search", monitors.Search.Handle)
			monitorsApi.PUT("/:id", monitors.Update.Handle)
			monitorsApi.DELETE("/:id", monitors.Delete.Handle)
		}

	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
