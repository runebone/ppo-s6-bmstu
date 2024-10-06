package main

import (
	"fmt"
	"log"
	"net/http"
	"user/internal/adapter/database"
	"user/internal/adapter/logger"
	sqlxRepo "user/internal/adapter/repository/sqlx"
	api "user/internal/api/v1"
	"user/internal/config"
	handler "user/internal/handler/v1"
	"user/internal/middleware"
	usecase "user/internal/usecase/v1"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Println("Error reading config (config.toml)")
	}

	db, err := database.NewPostgresDB(config.User.Postgres)
	if err != nil {
		log.Fatalf("Couldn't connect to database, exiting: %w", err)
	}

	logger := logger.NewZapLogger(config.User.Log)

	repo := sqlxRepo.NewSQLXUserRepository(db)
	uc := usecase.NewUserUseCase(repo)

	userHandler := handler.NewUserHandler(uc)
	router := mux.NewRouter()
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	router.Use(loggingMiddleware.Middleware)
	api.InitializeV1Routes(router, userHandler)

	localPort := fmt.Sprintf("%d", config.User.LocalPort)
	exposedPort := fmt.Sprintf("%d", config.User.ExposedPort)

	log.Printf("Starting server on :%s\n", exposedPort)
	http.ListenAndServe(":"+localPort, router)
}
