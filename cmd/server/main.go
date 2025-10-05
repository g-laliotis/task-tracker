package main

import (
	"github.com/g-laliotis/task-tracker/internal/handler"
	"github.com/g-laliotis/task-tracker/internal/model"
	"github.com/g-laliotis/task-tracker/internal/repository"
	"github.com/g-laliotis/task-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1️⃣ Connect to SQLite
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 2️⃣ Auto-migrate the Task model
	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		panic(err)
	}

	// 3️⃣ Initialize repository, service, and handler
	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	// 4️⃣ Setup Gin router and register routes
	r := gin.Default()
	h.RegisterRoutes(r)

	// 5️⃣ Start server
	r.Run(":8080")
}
