package persistence

import (
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(d *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: d,
	}
}

func (r TaskRepository) ListByUserID(id int) ([]model.Task, error) {
	var tasks []model.Task

	if res := r.db.Where("user_id = ?", id).Find(&tasks); res.Error != nil {
		return nil, hizumi.InternalServerError{Message: res.Error.Error()}
	} else if res.RowsAffected <= 0 {
		return nil, hizumi.NotFound{Message: gorm.ErrRecordNotFound.Error()}
	}

	return tasks, nil
}
