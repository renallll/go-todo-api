package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

// Data sementara (in-memory)
var tasks = []models.Task{
	{
		ID:        1,
		Title:     "Belajar Golang",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Belajar Gin",
		Completed: true,
	},
}

var nextID = 3

// GET /tasks
func GetTasks(c *gin.Context) {

	// Buat salinan slice agar data asli tidak berubah
	filteredTasks := make([]models.Task, len(tasks))
	copy(filteredTasks, tasks)

	// Ambil query parameter
	completed := c.Query("completed")

	// Filter jika ada query parameter
	if completed != "" {

		var result []models.Task

		isCompleted := completed == "true"

		for _, task := range filteredTasks {
			if task.Completed == isCompleted {
				result = append(result, task)
			}
		}

		filteredTasks = result
	}

	// Urutkan ID terbesar ke terkecil
	sort.Slice(filteredTasks, func(i, j int) bool {
		return filteredTasks[i].ID > filteredTasks[j].ID
	})

	c.JSON(http.StatusOK, filteredTasks)
}

// GET /tasks/:id
func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task tidak ditemukan",
	})
}

// POST /tasks
func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask.ID = nextID
	nextID++

	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updated models.Task

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			updated.ID = id
			tasks[i] = updated

			c.JSON(http.StatusOK, updated)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task tidak ditemukan",
	})
}

// DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "Task berhasil dihapus",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task tidak ditemukan",
	})
}
