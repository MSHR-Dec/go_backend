package repository

import (
	"github.com/MSHR-Dec/go_backend/internal/domain/model"
)

type TaskRepository interface {
	ListByUserID(id int) ([]model.Task, error)
}
