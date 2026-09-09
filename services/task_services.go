package services

import (
	"todo-api/models"
	"todo-api/repositories"
)

type TaskService struct {
	Repo repositories.TaskRepository
}

func (s *TaskService) GetTasks(userID uint) ([]models.Task, error) {
	return s.Repo.GetByUser(userID)
}

func (s *TaskService) GetTask(id string, userID uint) (*models.Task, error) {
	return s.Repo.GetByID(id, userID)
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(task *models.Task) error {
	return s.Repo.Delete(task)
}
