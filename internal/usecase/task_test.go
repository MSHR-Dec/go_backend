// mockgen -source ../domain/repository/task.go -destination mock_repository/mock_task_repository.go
package usecase_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	. "github.com/MSHR-Dec/go_backend/internal/usecase"
	"github.com/MSHR-Dec/go_backend/internal/usecase/mock_repository"
)

var _ = Describe("Task Interactor", func() {
	var (
		successID int

		failedID int

		mockCtrl           *gomock.Controller
		mockTaskRepository *mock_repository.MockTaskRepository

		task *TaskInteractor
	)

	BeforeEach(func() {
		successID = 1

		failedID = 0

		mockCtrl = gomock.NewController(GinkgoT())
		mockTaskRepository = mock_repository.NewMockTaskRepository(mockCtrl)
	})

	Describe("Test ListByUserID", func() {
		Context("The case that is valid user id", func() {
			It("should return task list", func() {
				input := ListByUserIDInput{
					ID: successID,
				}
				expected := []ListByUserIDOutput{
					{
						Name: "sample",
					},
				}

				mockTaskRepository.EXPECT().ListByUserID(1).Return([]model.Task{
					{
						ID: 1,
						Name: "sample",
						Summary: "A sample task for admin",
						UserID: 1,
					},
				}, nil)

				task = NewTaskInteractor(mockTaskRepository)
				output, err := task.ListByUserID(input)

				Expect(output).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})

		Context("The case that is invalid user id", func() {
			It("should return error", func() {
				input := ListByUserIDInput{
					ID: failedID,
				}
				mockTaskRepository.EXPECT().ListByUserID(0).Return([]model.Task{}, gorm.ErrRecordNotFound)

				task = NewTaskInteractor(mockTaskRepository)
				output, err := task.ListByUserID(input)

				Expect(output).To(BeNil())
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
			})
		})
	})
})
