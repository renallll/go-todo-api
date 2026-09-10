package services

import (
	"testing"
	"todo-api/models"
)

type MockTaskRepository struct {
	tasks []models.Task
}

func (m *MockTaskRepository) GetByUser(userID uint) ([]models.Task, error) {
	return m.tasks, nil
}

func (m *MockTaskRepository) GetByID(id string, userID uint) (*models.Task, error) {
	return &m.tasks[0], nil
}

func (m *MockTaskRepository) Create(task *models.Task) error {
	m.tasks = append(m.tasks, *task)
	return nil
}

func (m *MockTaskRepository) Update(task *models.Task) error {
	return nil
}

func (m *MockTaskRepository) Delete(task *models.Task) error {
	return nil
}

func TestGetTasks(t *testing.T) {

	mockRepo := &MockTaskRepository{
		tasks: []models.Task{
			{Title: "Belajar Go"},
			{Title: "Belajar JWT"},
		},
	}

	service := NewTaskService(mockRepo)

	tasks, err := service.GetTasks(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "Belajar Go" {
		t.Errorf("unexpected title")
	}
}

func TestCreateTask(t *testing.T) {

	mockRepo := &MockTaskRepository{}

	service := NewTaskService(mockRepo)

	task := &models.Task{
		Title: "Testing Golang",
	}

	err := service.CreateTask(task)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(mockRepo.tasks) != 1 {
		t.Errorf("expected 1 task")
	}
}
