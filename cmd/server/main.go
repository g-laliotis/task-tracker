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
	_ = godotenv.Load()

	// 2️⃣ Get environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL") // optional PostgreSQL
	var db *gorm.DB
	var err error

	// 3️⃣ Connect to database
	if dbURL != "" {
		fmt.Println("🔗 Connecting to PostgreSQL...")
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	} else {
		fmt.Println("💾 Using local SQLite database...")
		db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// 4️⃣ Auto-migrate the Task model
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	// 5️⃣ Initialize repository, service, and handler
	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	// 6️⃣ Setup Gin router and register routes
	r := gin.Default()
	h.RegisterRoutes(r)

	// 7️⃣ Start server (Render requires dynamic port binding)
	fmt.Println("🚀 Server running on port " + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
