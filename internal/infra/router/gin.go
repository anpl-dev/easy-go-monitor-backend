package router

import (
	"go-monitor-tool/internal/api/middleware"
	"go-monitor-tool/internal/infra/jwt"
	monitorHandler "go-monitor-tool/internal/monitor/adapter/handler"
	userHandler "go-monitor-tool/internal/user/adapter/handler"
	"github.com/gin-gonic/gin"
)

type (
	UserHandlers struct {
		Create   *userHandler.CreateUserHandler
		FindByID *userHandler.FindUserByIDHandler
		Search   *userHandler.SearchUsersHandler
		Update   *userHandler.UpdateUserHandler
		Delete   *userHandler.DeleteUserHandler
		Login    *userHandler.LoginUserHandler
	}

	MonitorHandlers struct {
		Create   *monitorHandler.CreateMonitorHandler
		FindByID *monitorHandler.FindMonitorByIDHandler
		Search   *monitorHandler.SearchMonitorsHandler
		Update   *monitorHandler.UpdateMonitorHandler
		Delete   *monitorHandler.DeleteMonitorHandler
	}
)

func NewGinRouter(users UserHandlers, monitors MonitorHandlers, jwtService jwt.JWTService) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api/v1")

	api.POST("/login", users.Login.Handle)
	api.POST("/users", users.Create.Handle)

	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleWare(jwtService))
	{
		usersApi := auth.Group("/users")
		{
			usersApi.GET("/:id", users.FindByID.Handle)
			usersApi.GET("/search", users.Search.Handle)
			usersApi.PUT("/:id", users.Update.Handle)
			usersApi.DELETE("/:id", users.Delete.Handle)
		}

		monitorsApi := auth.Group("/monitors")
		{
			monitorsApi.POST("", monitors.Create.Handle)
			monitorsApi.GET("/:id", monitors.FindByID.Handle)
			monitorsApi.GET("/search", monitors.Search.Handle)
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
