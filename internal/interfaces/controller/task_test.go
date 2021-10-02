// mockgen -source ../../usecase/task.go -destination mock_usecase/mock_task_usecase.go
package controller_test

import (
	"encoding/json"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/interfaces/controller"
	"github.com/MSHR-Dec/go_backend/internal/interfaces/controller/mock_usecase"
	"github.com/MSHR-Dec/go_backend/internal/usecase"
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"
)

var _ = Describe("Task Controller", func() {
	var (
		mockCtrl        *gomock.Controller
		mockTaskUsecase *mock_usecase.MockTaskUsecase

		task *controller.TaskController

		w *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockTaskUsecase = mock_usecase.NewMockTaskUsecase(mockCtrl)
	})

	Describe("Test ListTasksByUserID", func() {
		Context("The case that is given valid user id", func() {
			It("should returns status ok and list of task", func() {
				input := usecase.ListByUserIDInput{
					ID: 1,
				}
				output := []usecase.ListByUserIDOutput{
					{
						Name: "sample",
					},
				}
				mockTaskUsecase.EXPECT().ListByUserID(input).Return(output, nil)

				task = controller.NewTaskController(mockTaskUsecase)

				w = httptest.NewRecorder()

				expected, _ := json.Marshal(output)

				c, _ := gin.CreateTestContext(w)
				c.Params = []gin.Param{
					{
						Key: "userID",
						Value: "1",
					},
				}
				c.Request, _ = http.NewRequest("GET", "/tasks/1", nil)

				task.ListTasksByUserID(c)
				Expect(w.Code).To(Equal(200))
				Expect(w.Body.Bytes()).To(Equal(expected))
			})
		})

		Context("The case that is given invalid user id", func() {
			It("should returns status of not found", func() {
				input := usecase.ListByUserIDInput{
					ID: 0,
				}
				output := map[string]interface{}{
					"message": "record not found",
				}
				mockTaskUsecase.EXPECT().ListByUserID(input).Return(nil, hizumi.NotFound{Message: gorm.ErrRecordNotFound.Error()})

				task = controller.NewTaskController(mockTaskUsecase)

				w = httptest.NewRecorder()

				expected, _ := json.Marshal(output)

				c, _ := gin.CreateTestContext(w)
				c.Params = []gin.Param{
					{
						Key: "userID",
						Value: "0",
					},
				}
				c.Request, _ = http.NewRequest("GET", "/tasks/0", nil)

				task.ListTasksByUserID(c)
				Expect(w.Code).To(Equal(404))
				Expect(w.Body.Bytes()).To(Equal(expected))
			})
		})
	})
})
