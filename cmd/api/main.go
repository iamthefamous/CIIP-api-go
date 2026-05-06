package main

import (
	"log"
	"os"

	"github.com/iamthefamous/CIIP-api-go/internal/db"
	"github.com/iamthefamous/CIIP-api-go/internal/handler"
	"github.com/iamthefamous/CIIP-api-go/internal/repository"
	"github.com/iamthefamous/CIIP-api-go/internal/router"
	"github.com/iamthefamous/CIIP-api-go/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	database := db.NewPostgres()
	postRepo := repository.NewPostRepository(database)
	userRepo := repository.NewUserRepository(database)

	postService := service.NewPostService(postRepo)
	authService := service.NewAuthService(userRepo, jwtSecret)

	postHandler := handler.NewPostHandler(postService)
	authHandler := handler.NewAuthHandler(authService)

	r := router.NewRouter(postHandler, authHandler, jwtSecret)
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
