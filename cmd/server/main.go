package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"users/internal/handler"
	"users/internal/models"
	"users/internal/repository"
	"users/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	_ = db //
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHandler(userService)
	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetById)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	log.Fatal(router.Run(":8080"))

}
