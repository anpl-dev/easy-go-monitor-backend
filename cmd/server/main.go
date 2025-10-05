package main

import (
	"go-monitor-tool/internal/infra/database"
	"go-monitor-tool/internal/infra/router"
	monitorHandler "go-monitor-tool/internal/monitor/adapter/handler"
	monitorPresenter "go-monitor-tool/internal/monitor/adapter/presenter"
	monitorRepo "go-monitor-tool/internal/monitor/adapter/repository"
	monitorUC "go-monitor-tool/internal/monitor/usecase"
	userHandler "go-monitor-tool/internal/user/adapter/handler"
	userPresenter "go-monitor-tool/internal/user/adapter/presenter"
	userRepo "go-monitor-tool/internal/user/adapter/repository"
	userUC "go-monitor-tool/internal/user/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env found")
	}
}

func main() {
	// TODO: Middleware: auth
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
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	// --- Repository ---
	userRepo := userRepo.NewUserPostgresRepository(db)
	monitorRepo := monitorRepo.NewMonitorPostgresRepository(db)

	// --- Presenter ---
	createUserPresenter := userPresenter.NewCreateUserPresenter()
	findUserByIDPresenter := userPresenter.NewFindUserByIDPresenter()
	searchUsersPresenter := userPresenter.NewSearchUsersPresenter()
	updateUserPresenter := userPresenter.NewUpdateUserPresenter()

	createMonitorPresenter := monitorPresenter.NewCreateMonitorPresenter()
	findMonitorByIDPresenter := monitorPresenter.NewFindMonitorByIDPresenter()
	searchMonitorsPresenter := monitorPresenter.NewSearchMonitorsPresenter()
	updateMonitorPresenter := monitorPresenter.NewUpdateMonitorPresenter()

	// --- UseCase ---
	createUserUC := userUC.NewCreateUserInteractor(userRepo, createUserPresenter)
	findUserByIDUC := userUC.NewFindUserByIDInteractor(userRepo, findUserByIDPresenter)
	searchUsersUC := userUC.NewSearchUsersInteractor(userRepo, searchUsersPresenter)
	updateUserUC := userUC.NewUpdateUserInteractor(userRepo, updateUserPresenter)
	deleteUserUC := userUC.NewDeleteUserInteractor(userRepo)

	createMonitorUC := monitorUC.NewCreateMonitorInteractor(monitorRepo, createMonitorPresenter)
	findMointorByIDUC := monitorUC.NewFindMonitorByIDInteractor(monitorRepo, findMonitorByIDPresenter)
	searchMonitorsUC := monitorUC.NewSearchMonitorsInteractor(monitorRepo, searchMonitorsPresenter)
	updateMonitorUC := monitorUC.NewUpdateMonitorInteractor(monitorRepo, updateMonitorPresenter)
	deleteMonitorUC := monitorUC.NewDeleteMonitorInteractor(monitorRepo)

	// --- Handler ---
	userHandlers := router.UserHandlers{
		Create:   userHandler.NewCreateUserHandler(createUserUC),
		FindByID: userHandler.NewFindUserByIDHandler(findUserByIDUC),
		Search:   userHandler.NewSearchUsersHandler(searchUsersUC),
		Update:   userHandler.NewUpdateUserHandler(updateUserUC),
		Delete:   userHandler.NewDeleteUserHandler(deleteUserUC),
	}

	monitorHandlers := router.MonitorHandlers{
		Create:   monitorHandler.NewCreateMonitorHandler(createMonitorUC),
		FindByID: monitorHandler.NewFindMonitorByIDHandler(findMointorByIDUC),
		Search:   monitorHandler.NewSearchMonitorsHandler(searchMonitorsUC),
		Update:   monitorHandler.NewUpdateMonitorHandler(updateMonitorUC),
		Delete:   monitorHandler.NewDeleteMonitorHandler(deleteMonitorUC),
	}

	// --- Router ---
	r := router.NewGinRouter(userHandlers, monitorHandlers)

	// --- Run Server ---
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("falled to start server: %v", err)
	}
}
