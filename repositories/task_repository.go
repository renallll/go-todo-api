package repositories

import (
	"todo-api/database"
	"todo-api/models"
)

type TaskRepository struct{}

func (r *TaskRepository) GetByUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task

	err := database.DB.
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&tasks).Error

	return tasks, err
}

func (r *TaskRepository) GetByID(id string, userID uint) (*models.Task, error) {
	var task models.Task

	err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error

	return &task, err
}

func (r *TaskRepository) Create(task *models.Task) error {
	return database.DB.Create(task).Error
}

func (r *TaskRepository) Update(task *models.Task) error {
	return database.DB.Save(task).Error
}

func (r *TaskRepository) Delete(task *models.Task) error {
	return database.DB.Delete(task).Error
}
