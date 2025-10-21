package router

import (
	"go-monitor-tool/internal/api/middleware"
	"go-monitor-tool/internal/infra/jwt"
	monitorController "go-monitor-tool/internal/monitor/adapter/controller"
	userController "go-monitor-tool/internal/user/adapter/controller"

	"github.com/gin-gonic/gin"
)

type (
	UserControllers struct {
		Create   *userController.CreateUserController
		FindByID *userController.FindUserByIDController
		Search   *userController.SearchUsersController
		Update   *userController.UpdateUserController
		Delete   *userController.DeleteUserController
		Login    *userController.LoginUserController
	}

	MonitorControllers struct {
		Create   *monitorController.CreateMonitorController
		FindByID *monitorController.FindMonitorByIDController
		Search   *monitorController.SearchMonitorsController
		Update   *monitorController.UpdateMonitorController
		Delete   *monitorController.DeleteMonitorController
	}
)

func NewGinRouter(users UserControllers, monitors MonitorControllers, jwtService jwt.JWTService) *gin.Engine {
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
