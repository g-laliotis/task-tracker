package main

import (
	"fmt"
	"log"
	"os"

	"github.com/g-laliotis/task-tracker/internal/handler"
	"github.com/g-laliotis/task-tracker/internal/model"
	"github.com/g-laliotis/task-tracker/internal/repository"
	"github.com/g-laliotis/task-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load .env
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	dbURL := os.Getenv("DATABASE_URL")

	var db *gorm.DB
	var err error

	if dbURL != "" {
		fmt.Println("🔗 Connecting to PostgreSQL...")
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("💾 Using SQLite...")
		db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := db.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userSvc)

	taskRepo := repository.NewTaskRepository(db)
	taskSvc := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskSvc)

	r := gin.Default()

	// Auth routes
	authHandler.RegisterRoutes(r)

	// Task routes (JWT-protected)
	taskHandler.RegisterRoutes(r)

	fmt.Println("🚀 Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
