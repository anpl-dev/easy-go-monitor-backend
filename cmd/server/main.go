package main

import (
	"easy-go-monitor/internal/infra/database"
	"easy-go-monitor/internal/infra/jwt"
	"easy-go-monitor/internal/infra/router"
	monitorController "easy-go-monitor/internal/monitor/adapter/controller"
	monitorPresenter "easy-go-monitor/internal/monitor/adapter/presenter"
	monitorRepo "easy-go-monitor/internal/monitor/adapter/repository"
	monitorUC "easy-go-monitor/internal/monitor/usecase"
	userController "easy-go-monitor/internal/user/adapter/controller"
	userPresenter "easy-go-monitor/internal/user/adapter/presenter"
	userRepo "easy-go-monitor/internal/user/adapter/repository"
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

	// --- JWT Service ---
	expHour, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOUR"))
	jwtService := jwt.NewService(os.Getenv("JWT_SECRET"), time.Duration(expHour)*time.Hour)

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
	loginUserUC := userUC.NewLoginUserInteractor(userRepo, jwtService)

	createMonitorUC := monitorUC.NewCreateMonitorInteractor(monitorRepo, createMonitorPresenter)
	findMointorByIDUC := monitorUC.NewFindMonitorByIDInteractor(monitorRepo, findMonitorByIDPresenter)
	searchMonitorsUC := monitorUC.NewSearchMonitorsInteractor(monitorRepo, searchMonitorsPresenter)
	updateMonitorUC := monitorUC.NewUpdateMonitorInteractor(monitorRepo, updateMonitorPresenter)
	deleteMonitorUC := monitorUC.NewDeleteMonitorInteractor(monitorRepo)

	// --- Controller ---
	userControllers := router.UserControllers{
		Create:   userController.NewCreateUserController(createUserUC),
		FindByID: userController.NewFindUserByIDController(findUserByIDUC),
		Search:   userController.NewSearchUsersController(searchUsersUC),
		Update:   userController.NewUpdateUserController(updateUserUC),
		Delete:   userController.NewDeleteUserController(deleteUserUC),
		Login:    userController.NewLoginUserController(loginUserUC),
	}

	monitorControllers := router.MonitorControllers{
		Create:   monitorController.NewCreateMonitorController(createMonitorUC),
		FindByID: monitorController.NewFindMonitorByIDController(findMointorByIDUC),
		Search:   monitorController.NewSearchMonitorsController(searchMonitorsUC),
		Update:   monitorController.NewUpdateMonitorController(updateMonitorUC),
		Delete:   monitorController.NewDeleteMonitorController(deleteMonitorUC),
	}

	// --- Router ---
	r := router.NewGinRouter(userControllers, monitorControllers, jwtService)

	// --- Run Server ---
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("falled to start server: %v", err)
	}
}
