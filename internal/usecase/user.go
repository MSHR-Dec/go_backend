package usecase

import (
	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	"github.com/MSHR-Dec/go_backend/internal/domain/repository"
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"
)

type UserUsecase interface {
	SignIn(input SignInInput) (SignInOutput, error)
}

type UserInteractor struct {
	repository repository.UserRepository
}

func NewUserInteractor(r repository.UserRepository) *UserInteractor {
	return &UserInteractor{
		repository: r,
	}
}

type SignInInput struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignInOutput struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

func toSignInOutput(user model.User) SignInOutput {
	return SignInOutput{
		ID: user.ID,
		Name: user.Name,
		Nickname: user.Nickname,
	}
}

func (i *UserInteractor) SignIn(input SignInInput) (SignInOutput, error) {
	user, err := i.repository.GetByName(input.Name)
	if err != nil {
		return SignInOutput{}, err
	}

	if user.IsCollectPassword(input.Password) {
		return SignInOutput{}, hizumi.BadRequest{Message: "invalid password"}
	}

	return toSignInOutput(user), nil
}
