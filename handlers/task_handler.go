package handlers

import (
	"net/http"

	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

// GET /tasks
func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var tasks []models.Task

	if err := database.DB.
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&tasks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil task",
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GET /tasks/:id
func GetTaskByID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var task models.Task

	if err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error; err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Task tidak ditemukan atau bukan milik Anda",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// POST /tasks
func CreateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.UserID = userID.(uint)

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat task",
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var task models.Task

	if err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error; err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Task tidak ditemukan atau bukan milik Anda",
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
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var task models.Task

	if err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error; err != nil {

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Task tidak ditemukan atau bukan milik Anda",
		})
		return
	}

	database.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task berhasil dihapus",
	})
}
