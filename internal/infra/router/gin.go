package router

import (
	"easy-go-monitor/internal/api/middleware"
	"easy-go-monitor/internal/infra/jwt"
	"easy-go-monitor/internal/infra/logger"
	monitorController "easy-go-monitor/internal/monitor/adapter/controller"
	runnerController "easy-go-monitor/internal/runner/adapter/controller"
	userController "easy-go-monitor/internal/user/adapter/controller"

	"github.com/gin-gonic/gin"
)

type (
	UserControllers struct {
		Create   *userController.CreateUserController
		FindByID *userController.FindUserByIDController
		Update   *userController.UpdateUserController
		Delete   *userController.DeleteUserController
		Login    *userController.LoginUserController
	}

	MonitorControllers struct {
		Create   *monitorController.CreateMonitorController
		FindByID *monitorController.FindMonitorByIDController
		FindAll  *monitorController.FindAllMonitorsController
		Update   *monitorController.UpdateMonitorController
		Delete   *monitorController.DeleteMonitorController

		SetEnabled *monitorController.SetEnabledMonitorController
	}

	RunnerControllers struct {
		Create   *runnerController.CreateRunnerController
		FindByID *runnerController.FindRunnerByIDController
		FindAll  *runnerController.FindAllRunnersController
		Update   *runnerController.UpdateRunnerController
		Delete   *runnerController.DeleteRunnerController
		Execute  *runnerController.ExecuteRunnerController
		History  *runnerController.FindRunnerHistoriesController
		Search   *runnerController.SearchRunnerHistoriesController
	}
)

func NewGinRouter(
	users UserControllers,
	monitors MonitorControllers,
	runners RunnerControllers,
	jwtService jwt.JWTService,
	log *logger.Logger,
) *gin.Engine {
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.LoggerMiddleWare(log),
	)

	api := r.Group("/api/v1")

	// login or create users are not jwt authenticated
	api.POST("/login", users.Login.Handle)
	api.POST("/users", users.Create.Handle)

	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleWare(jwtService))
	{
		usersApi := auth.Group("/users")
		{
			usersApi.GET("/:id", users.FindByID.Handle)
			usersApi.PUT("/:id", users.Update.Handle)
			usersApi.DELETE("/:id", users.Delete.Handle)
		}

		monitorsApi := auth.Group("/monitors")
		{
			monitorsApi.POST("", monitors.Create.Handle)
			monitorsApi.GET("/:id", monitors.FindByID.Handle)
			monitorsApi.GET("", monitors.FindAll.Handle)
			monitorsApi.PUT("/:id", monitors.Update.Handle)
			monitorsApi.PATCH("/:id/enabled", monitors.SetEnabled.Handle)
			monitorsApi.DELETE("/:id", monitors.Delete.Handle)
		}

		runnersApi := auth.Group("/runners")
		{
			runnersApi.POST("", runners.Create.Handle)
			runnersApi.GET("/:id", runners.FindByID.Handle)
			runnersApi.GET("", runners.FindAll.Handle)
			runnersApi.PUT("/:id", runners.Update.Handle)
			runnersApi.DELETE("/:id", runners.Delete.Handle)
			runnersApi.POST("/:id/execute", runners.Execute.Handle)
			runnersApi.GET("/:id/histories", runners.History.Handle)
			runnersApi.GET("/histories", runners.Search.Handle)
		}
	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
