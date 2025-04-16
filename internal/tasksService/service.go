package tasksService

import (
	"awesomeProject/internal/models"
)

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask создает новую задачу
func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.CreateTask(task)
}

// GetAllTasks возвращает все задачи из базы данных
func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

// GetTasksByUserID возвращает все задачи пользователя
func (s *TaskService) GetTasksByUserID(userID uint) ([]models.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

// UpdateTaskByID обновляет задачу по её ID
func (s *TaskService) UpdateTaskByID(id uint, task models.Task) (models.Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

// DeleteTaskByID удаляет задачу по её ID
func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
