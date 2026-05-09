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
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseAnonKey := os.Getenv("SUPABASE_ANON_KEY")
	supabaseJWTSecret := os.Getenv("SUPABASE_JWT_SECRET")

	database := db.NewPostgres()

	postRepo := repository.NewPostRepository(database)
	userRepo := repository.NewUserRepository(database)
	universityRepo := repository.NewUniversityRepository(database)
	facultyRepo := repository.NewFacultyRepository(database)
	programRepo := repository.NewProgramRepository(database)

	postService := service.NewPostService(postRepo)
	authService := service.NewAuthService(userRepo, supabaseURL, supabaseAnonKey)
	universityService := service.NewUniversityService(universityRepo)
	facultyService := service.NewFacultyService(facultyRepo)
	programService := service.NewProgramService(programRepo)

	postHandler := handler.NewPostHandler(postService)
	authHandler := handler.NewAuthHandler(authService)
	universityHandler := handler.NewUniversityHandler(universityService)
	facultyHandler := handler.NewFacultyHandler(facultyService)
	programHandler := handler.NewProgramHandler(programService)

	r := router.NewRouter(
		postHandler,
		authHandler,
		universityHandler,
		facultyHandler,
		programHandler,
		jwtSecret,
		supabaseJWTSecret,
		userRepo.GetProfileRoleByID,
	)
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
