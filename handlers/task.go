package handlers

import (
	"net/http"
	"gotasker/database"
	"gotasker/models"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
}


func GetTasks(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	
	var tasks []models.Task
	if err := database.DB.Where("user_id = ?", user.ID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	taskID := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status := models.TaskStatus(req.Status)
	if status != "" && status != models.StatusTodo && status != models.StatusInProgress && status != models.StatusDone {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	if status == "" {
		status = models.StatusTodo
	}

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
		Priority:    req.Priority,
		UserID:      user.ID,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	taskID := c.Param("id")

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Status != "" {
		status := models.TaskStatus(req.Status)
		if status != models.StatusTodo && status != models.StatusInProgress && status != models.StatusDone {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
			return
		}
		task.Status = status
	}
	if req.Priority > 0 {
		task.Priority = req.Priority
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	taskID := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, user.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
