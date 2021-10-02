package controller

import (
	"net/http"
	"strconv"

	"github.com/MSHR-Dec/go_backend/internal/usecase"
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"
)

type TaskController struct {
	usecase usecase.TaskUsecase
}

func NewTaskController(u usecase.TaskUsecase) *TaskController {
	return &TaskController{
		usecase: u,
	}
}

func (c TaskController) ListTasksByUserID(ctx Context) {
	var input usecase.ListByUserIDInput
	input.ID, _ = strconv.Atoi(ctx.Param("userID"))

	tasks, err := c.usecase.ListByUserID(input)

	if err != nil {
		switch err.(type) {
		case hizumi.NotFound:
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"message": err.Error()})
			return
		case hizumi.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, tasks)
}
