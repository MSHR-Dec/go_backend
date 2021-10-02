// +build wireinject

// cd path/to/internal/infrastructure && wire
package infrastructure

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/repository"
	"github.com/MSHR-Dec/go_backend/internal/interfaces/controller"
	"github.com/MSHR-Dec/go_backend/internal/interfaces/persistence"
	"github.com/MSHR-Dec/go_backend/internal/usecase"
)

func initializeUser(db *gorm.DB) *controller.UserController {
	wire.Build(
		controller.NewUserController,
		usecase.NewUserInteractor,
		persistence.NewUserRepository,
		wire.Bind(new(usecase.UserUsecase), new(*usecase.UserInteractor)),
		wire.Bind(new(repository.UserRepository), new(*persistence.UserRepository)),
	)

	return &controller.UserController{}
}

func initializeTask(db *gorm.DB) *controller.TaskController {
	wire.Build(
		controller.NewTaskController,
		usecase.NewTaskInteractor,
		persistence.NewTaskRepository,
		wire.Bind(new(usecase.TaskUsecase), new(*usecase.TaskInteractor)),
		wire.Bind(new(repository.TaskRepository), new(*persistence.TaskRepository)),
	)

	return &controller.TaskController{}
}
