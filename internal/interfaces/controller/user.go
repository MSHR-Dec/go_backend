package controller

import (
	"net/http"

	"github.com/MSHR-Dec/go_backend/internal/usecase"
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"
)

type UserController struct {
	usecase usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) *UserController {
	return &UserController{
		usecase: u,
	}
}

func (c *UserController) SignIn(ctx Context) {
	var input usecase.SignInInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": "invalid parameter"})
		return
	}

	user, err := c.usecase.SignIn(input)

	if err != nil {
		switch err.(type) {
		case hizumi.NotFound:
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"message": err.Error()})
			return
		case hizumi.BadRequest:
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			return
		case hizumi.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, user)
}
