package main

import (
	"go-monitor-tool/internal/adapter/handler"
	"go-monitor-tool/internal/adapter/presenter"
	"go-monitor-tool/internal/adapter/repository"
	"go-monitor-tool/internal/infra/database"
	"go-monitor-tool/internal/infra/router"
	"go-monitor-tool/internal/usecase"
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
	userRepo := repository.NewUserPostgresRepository(db)
	monitorRepo := repository.NewMonitorPostgresRepository(db)

	// --- Presenter ---
	createUserPresenter := presenter.NewCreateUserPresenter()
	findUserByIDPresenter := presenter.NewFindUserByIDPresenter()
	findUserByEmailPresenter := presenter.NewFindUserByEmailPresenter()
	createMonitorPresenter := presenter.NewCreateMonitorPresenter()

	// --- UseCase ---
	createUserUC := usecase.NewCreateUser(userRepo, createUserPresenter)
	findUserByIDUC := usecase.NewFindUserByID(userRepo, findUserByIDPresenter)
	findUserByEmailUC := usecase.NewFindUserByEmail(userRepo, findUserByEmailPresenter)

	createMonitorUC := usecase.NewCreateMonitor(monitorRepo, createMonitorPresenter)

	// --- Handler ---
	userHandlers := router.UserHandlers{
		Create:      handler.NewCreateUserHandler(createUserUC),
		FindByID:    handler.NewFindUserByIDHandler(findUserByIDUC),
		FindByEmail: handler.NewFindUserByEmailHandler(findUserByEmailUC),
	}

	monitorHandlers := router.MonitorHandlers{
		Create: handler.NewCreateMonitorHandler(createMonitorUC),
	}

	// --- Router ---
	r := router.NewGinRouter(userHandlers, monitorHandlers)

	// --- Run Server ---
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("falled to start server: %v", err)
	}
}
