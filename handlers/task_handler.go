package handlers

import (
	"net/http"

	"todo-api/models"
	"todo-api/repositories"
	"todo-api/services"

	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	Title     string `json:"title" example:"Belajar Go"`
	Completed bool   `json:"completed" example:"false"`
}

type UpdateTaskInput struct {
	Title     string `json:"title" example:"Belajar Swagger"`
	Completed bool   `json:"completed" example:"true"`
}

type TaskResponse struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"Belajar Go"`
	Completed bool   `json:"completed" example:"false"`
	UserID    uint   `json:"user_id" example:"1"`
}

// GetTasks godoc
// @Summary Get semua task
// @Description Mengambil semua task milik user yang sedang login
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Success 200 {array} TaskResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")

	repo := repositories.TaskRepository{}
	service := services.NewTaskService(&repo)

	tasks, err := service.GetTasks(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil task",
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// @Summary Get task berdasarkan ID
// @Description Mengambil satu task berdasarkan ID milik user yang sedang login
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} TaskResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /tasks/{id} [get]
func GetTaskByID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	repo := repositories.TaskRepository{}
	service := services.NewTaskService(&repo)

	task, err := service.GetTask(id, userID.(uint))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Task tidak ditemukan atau bukan milik Anda",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask godoc
// @Summary Membuat task baru
// @Description Membuat task baru untuk user yang sedang login
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body CreateTaskInput true "Task"
// @Success 201 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
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

	repo := repositories.TaskRepository{}
	service := services.NewTaskService(&repo)

	if err := service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat task",
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
// UpdateTask godoc
// @Summary Update task
// @Description Mengubah task milik user yang sedang login
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param body body UpdateTaskInput true "Task"
// @Success 200 {object} TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	repo := repositories.TaskRepository{}
	service := services.NewTaskService(&repo)

	task, err := service.GetTask(id, userID.(uint))
	if err != nil {
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

	if err := service.UpdateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengupdate task",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Hapus task
// @Description Menghapus task milik user yang sedang login
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	repo := repositories.TaskRepository{}
	service := services.NewTaskService(&repo)

	task, err := service.GetTask(id, userID.(uint))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Task tidak ditemukan atau bukan milik Anda",
		})
		return
	}

	if err := service.DeleteTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menghapus task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task berhasil dihapus",
	})
}
