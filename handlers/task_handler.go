package handlers

import (
	"net/http"

	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

// GET /tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task

	if err := database.DB.Order("id DESC").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data",
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GET /tasks/:id
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// POST /tasks
func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&task)

	c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task tidak ditemukan",
		})
		return
	}

	var input models.Task

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.Title = input.Title
	task.Completed = input.Completed

	database.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

// DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task tidak ditemukan",
		})
		return
	}

	database.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task berhasil dihapus",
	})
}
