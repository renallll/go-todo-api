package services

import (
	"todo-api/models"
)

type TaskRepository interface {
	GetByUser(userID uint) ([]models.Task, error)
	GetByID(id string, userID uint) (*models.Task, error)
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(task *models.Task) error
}

type TaskService struct {
	Repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
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
