// mockgen -source ../domain/repository/user.go -destination mock_repository/mock_user_repository.go
package usecase_test

import (
	"github.com/MSHR-Dec/go_backend/pkg/hizumi"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	. "github.com/MSHR-Dec/go_backend/internal/usecase"
	"github.com/MSHR-Dec/go_backend/internal/usecase/mock_repository"
)

var _ = Describe("User Interactor", func() {
	var (
		mockCtrl           *gomock.Controller
		mockUserRepository *mock_repository.MockUserRepository

		user *UserInteractor
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUserRepository = mock_repository.NewMockUserRepository(mockCtrl)
	})

	Describe("Test SignIn", func() {
		Context("The case that is valid user", func() {
			It("should return nil", func() {
				input := SignInInput{
					Name:     "admin",
					Password: "admin",
				}

				expected := SignInOutput{
					ID:       1,
					Name:     "admin",
					Nickname: "administrator",
				}

				mockUserRepository.EXPECT().GetByName("admin").Return(model.User{ID: 1, Name: "admin", Password: "admin", Nickname: "administrator"}, nil)

				user = NewUserInteractor(mockUserRepository)
				output, err := user.SignIn(input)

				Expect(output).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})

		Context("The case that is invalid user", func() {
			It("should return error", func() {
				input := SignInInput{
					Name: "dummy",
					Password: "admin",
				}

				mockUserRepository.EXPECT().GetByName("dummy").Return(model.User{}, gorm.ErrRecordNotFound)

				user = NewUserInteractor(mockUserRepository)
				output, err := user.SignIn(input)

				Expect(output).To(Equal(SignInOutput{}))
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
			})
		})

		Context("The case that is invalid password", func() {
			It("should return error", func() {
				input := SignInInput{
					Name:     "admin",
					Password: "dummy",
				}

				mockUserRepository.EXPECT().GetByName("admin").Return(model.User{ID: 1, Name: "admin", Password: "admin", Nickname: "administrator"}, nil)

				user = NewUserInteractor(mockUserRepository)
				output, err := user.SignIn(input)

				Expect(output).To(Equal(SignInOutput{}))
				Expect(err).To(Equal(hizumi.BadRequest{Message: "invalid password"}))
			})
		})
	})
})
