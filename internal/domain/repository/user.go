package repository

import (
	"github.com/MSHR-Dec/go_backend/internal/domain/model"
)

type UserRepository interface {
	GetByName(name string) (model.User, error)
}
