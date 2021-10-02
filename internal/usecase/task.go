package usecase

import (
	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	"github.com/MSHR-Dec/go_backend/internal/domain/repository"
)

type TaskUsecase interface {
	ListByUserID(input ListByUserIDInput) ([]ListByUserIDOutput, error)
}

type TaskInteractor struct {
	repository repository.TaskRepository
}

func NewTaskInteractor(r repository.TaskRepository) *TaskInteractor {
	return &TaskInteractor{
		repository: r,
	}
}

type ListByUserIDInput struct {
	ID int
}

type ListByUserIDOutput struct {
	Name string `json:"name"`
}

func toListByUserIDOutput(tasks []model.Task) []ListByUserIDOutput {
	var outputs []ListByUserIDOutput
	for _, task := range tasks {
		outputs = append(outputs, ListByUserIDOutput{Name: task.Name})
	}

	return outputs
}

func (i TaskInteractor) ListByUserID(input ListByUserIDInput) ([]ListByUserIDOutput, error) {
	tasks, err := i.repository.ListByUserID(input.ID)
	if err != nil {
		return nil, err
	}

	return toListByUserIDOutput(tasks), nil
}
