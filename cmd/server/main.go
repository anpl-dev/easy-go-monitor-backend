package main

import (
	"easy-go-monitor/internal/infra/database"
	"easy-go-monitor/internal/infra/jwt"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/infra/router"
	monitorController "easy-go-monitor/internal/monitor/adapter/controller"
	monitorPresenter "easy-go-monitor/internal/monitor/adapter/presenter"
	monitorRepository "easy-go-monitor/internal/monitor/adapter/repository"
	monitorUC "easy-go-monitor/internal/monitor/usecase"
	runnerController "easy-go-monitor/internal/runner/adapter/controller"
	runnerPresenter "easy-go-monitor/internal/runner/adapter/presenter"
	runnerRepository "easy-go-monitor/internal/runner/adapter/repository"
	runnerDomain "easy-go-monitor/internal/runner/domain"
	runnerUC "easy-go-monitor/internal/runner/usecase"
	userController "easy-go-monitor/internal/user/adapter/controller"
	userPresenter "easy-go-monitor/internal/user/adapter/presenter"
	userRepository "easy-go-monitor/internal/user/adapter/repository"
	userUC "easy-go-monitor/internal/user/usecase"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env found")
	}
}

func main() {
	// --- Logger ---
	appLogger := logger.NewLogger(os.Getenv("LOG_LEVEL"))
	appLogger.Info("Starting Easy-Go-Monitor backend...")

	// --- DB Config ---
	db_cfg := database.Config{
		Host:      os.Getenv("POSTGRES_HOST"),
		Port:      os.Getenv("POSTGRES_PORT"),
		User:      os.Getenv("POSTGRES_USER"),
		Passsword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:    os.Getenv("POSTGRES_DB"),
		SSLMode:   os.Getenv("POSTGRES_SSLMODE"),
	}

	db, err := database.NewPostgresDB(db_cfg)
	if err != nil {
		appLogger.Fatal("Failed to connect database", "error", err)
	}
	defer db.Close()

	appLogger.Info("Database connected successfully.")

	// --- JWT Service ---
	expHour, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOUR"))
	jwtService := jwt.NewService(os.Getenv("JWT_SECRET"), time.Duration(expHour)*time.Hour)

	// --- Repository ---
	userRepo := userRepository.NewUserPostgresRepository(db)
	monitorRepo := monitorRepository.NewMonitorPostgresRepository(db, appLogger)
	runnerRepo := runnerRepository.NewRunnerPostgresRepository(db)
	runnerHistoryRepo := runnerRepository.NewRunnerHistoryPostgresRepository(db, appLogger)

	// --- Presenter ---
	createUserPresenter := userPresenter.NewCreateUserPresenter()
	findUserByIDPresenter := userPresenter.NewFindUserByIDPresenter()
	updateUserPresenter := userPresenter.NewUpdateUserPresenter()

	createMonitorPresenter := monitorPresenter.NewCreateMonitorPresenter()
	findMonitorByIDPresenter := monitorPresenter.NewFindMonitorByIDPresenter()
	findAllMonitorsPresenter := monitorPresenter.NewFindAllMonitorsPresenter()
	updateMonitorPresenter := monitorPresenter.NewUpdateMonitorPresenter()
	setEnabledMonitorPresenter := monitorPresenter.NewSetEnabledMonitorPresenter()

	createRunnerPresenter := runnerPresenter.NewCreateRunnerPresenter()
	findRunnerByIDPresenter := runnerPresenter.NewFindRunnerByIDPresenter()
	findAllRunnersPresenter := runnerPresenter.NewFindAllRunnersPresenter()
	updateRunnerPresenter := runnerPresenter.NewUpdateRunnerPresenter()
	executeRunnerPresenter := runnerPresenter.NewExecuteRunnerPresenter()
	findRunnerHistoriesPresenter := runnerPresenter.NewFindRunnerHistoriesPresenter()
	searchRunnerHistoriesPresenter := runnerPresenter.NewSearchRunnerHistoriesPresenter()

	// --- Weboscket Notifier & RunnerService ---
	runnerService := runnerDomain.NewRunnerService(runnerRepo, monitorRepo, runnerHistoryRepo, appLogger)

	// --- UseCase ---
	createUserUC := userUC.NewCreateUserInteractor(userRepo, createUserPresenter)
	findUserByIDUC := userUC.NewFindUserByIDInteractor(userRepo, findUserByIDPresenter)
	updateUserUC := userUC.NewUpdateUserInteractor(userRepo, updateUserPresenter)
	deleteUserUC := userUC.NewDeleteUserInteractor(userRepo)
	loginUserUC := userUC.NewLoginUserInteractor(userRepo, jwtService)

	createMonitorUC := monitorUC.NewCreateMonitorInteractor(monitorRepo, createMonitorPresenter, appLogger)
	findMointorByIDUC := monitorUC.NewFindMonitorByIDInteractor(monitorRepo, findMonitorByIDPresenter)
	findAllMonitorsUC := monitorUC.NewFindAllMonitorsInteractor(monitorRepo, findAllMonitorsPresenter)
	updateMonitorUC := monitorUC.NewUpdateMonitorInteractor(monitorRepo, updateMonitorPresenter)
	deleteMonitorUC := monitorUC.NewDeleteMonitorInteractor(monitorRepo)
	setEnabledUC := monitorUC.NewSetEnabledMonitorInteractor(monitorRepo, setEnabledMonitorPresenter)

	createRunnerUC := runnerUC.NewCreateRunnerInteractor(runnerRepo, createRunnerPresenter)
	findRunnerUC := runnerUC.NewFindRunnerByIDInteractor(runnerRepo, findRunnerByIDPresenter)
	findAllRunnersUC := runnerUC.NewFindAllRunnersInteractor(runnerRepo, findAllRunnersPresenter)
	updateRunnerUC := runnerUC.NewUpdateRunnerInteractor(runnerRepo, updateRunnerPresenter)
	deleteRunnerUC := runnerUC.NewDeleteRunnerInteractor(runnerRepo)
	executeRunnerUC := runnerUC.NewExecuteRunnerInteractor(runnerService, executeRunnerPresenter, appLogger)
	findRunnerHistoriesUC := runnerUC.NewFindRunnerHistoriesInteractor(runnerHistoryRepo, findRunnerHistoriesPresenter)
	searchRunnerHistoriesUC := runnerUC.NewSearchRunnerHistoriesInteractor(runnerHistoryRepo, searchRunnerHistoriesPresenter, appLogger)

	// --- Controller ---
	userControllers := router.UserControllers{
		Create:   userController.NewCreateUserController(createUserUC),
		FindByID: userController.NewFindUserByIDController(findUserByIDUC),
		Update:   userController.NewUpdateUserController(updateUserUC),
		Delete:   userController.NewDeleteUserController(deleteUserUC),
		Login:    userController.NewLoginUserController(loginUserUC),
	}

	monitorControllers := router.MonitorControllers{
		Create:     monitorController.NewCreateMonitorController(createMonitorUC, appLogger),
		FindByID:   monitorController.NewFindMonitorByIDController(findMointorByIDUC),
		FindAll:    monitorController.NewFindAllMonitorsController(findAllMonitorsUC),
		Update:     monitorController.NewUpdateMonitorController(updateMonitorUC),
		Delete:     monitorController.NewDeleteMonitorController(deleteMonitorUC),
		SetEnabled: monitorController.NewSetEnabledMonitorController(setEnabledUC),
	}

	runnerControllers := router.RunnerControllers{
		Create:   runnerController.NewCreateRunnerController(createRunnerUC),
		FindByID: runnerController.NewFindRunnerByIDController(findRunnerUC),
		FindAll:  runnerController.NewFindAllRunnersController(findAllRunnersUC),
		Update:   runnerController.NewUpdateRunnerController(updateRunnerUC),
		Delete:   runnerController.NewDeleteRunnerController(deleteRunnerUC),
		Execute:  runnerController.NewExecuteRunnerController(executeRunnerUC, appLogger),
		History:  runnerController.NewFindRunnerHistoriesController(findRunnerHistoriesUC),
		Search:   runnerController.NewSearchRunnerHistoriesController(searchRunnerHistoriesUC),
	}

	// --- Gin Router ---
	r := router.NewGinRouter(userControllers, monitorControllers, runnerControllers, jwtService, appLogger)

	// --- Run Server ---
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := host + ":" + port
	appLogger.Info("Server starting", "addr", addr)
	if err := r.Run(addr); err != nil {
		appLogger.Fatal("Failed to start server", "error", err)
	}
}
