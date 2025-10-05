package main

import (
	"github.com/g-laliotis/task-tracker/internal/handler"
	"github.com/g-laliotis/task-tracker/internal/repository"
	"github.com/g-laliotis/task-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)
	r.Run(":8080")
}
