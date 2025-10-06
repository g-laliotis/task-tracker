package handler

import (
	"net/http"
	"strconv"

	"github.com/g-laliotis/task-tracker/internal/model"
	"github.com/g-laliotis/task-tracker/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/tasks", h.GetTasks)
	r.POST("/tasks", h.CreateTask)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)
}

// GET /tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// POST /tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input model.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if err := h.service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// PUT /tasks/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var input model.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	input.ID = uint(id)
	if err := h.service.Update(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

// DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
