package main

import (
	"log"

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

	database := db.NewPostgres()
	postRepo := repository.NewPostRepository(database)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := router.NewRouter(postHandler)
	log.Println("Starting server on :8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
