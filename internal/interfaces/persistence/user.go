package persistence

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(d *gorm.DB) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) GetByName(name string) (model.User, error) {
	var user model.User

	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return model.User{}, hizumi.NotFound{Message: err.Error()}
		} else {
			log.Println(err)
			return model.User{}, hizumi.InternalServerError{Message: err.Error()}
		}
	}

	return user, nil
}
