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
	// 1️⃣ Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	// 2️⃣ Read environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	dbURL := os.Getenv("DATABASE_URL")

	// 3️⃣ Connect to the database
	var db *gorm.DB
	if dbURL != "" {
		fmt.Println("🔗 Connecting to PostgreSQL...")
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to PostgreSQL:", err)
		}
	} else {
		fmt.Println("💾 Using local SQLite database...")
		db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to SQLite:", err)
		}
	}

	// 4️⃣ Auto-migrate models
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		log.Fatal("❌ Failed to migrate Task model:", err)
	}

	// 5️⃣ Initialize repository, service, handler
	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	// 6️⃣ Setup Gin router and register routes
	r := gin.Default()
	h.RegisterRoutes(r)

	// 7️⃣ Start server
	fmt.Println("🚀 Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
