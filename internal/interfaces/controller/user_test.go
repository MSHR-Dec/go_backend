// mockgen -source ../../usecase/user.go -destination mock_usecase/mock_use_usecase.go
package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/MSHR-Dec/go_backend/internal/interfaces/controller"
	"github.com/MSHR-Dec/go_backend/internal/interfaces/controller/mock_usecase"
	"github.com/MSHR-Dec/go_backend/internal/usecase"
)

var _ = Describe("User Controller", func() {
	var (
		output usecase.SignInOutput

		mockCtrl        *gomock.Controller
		mockUserUsecase *mock_usecase.MockUserUsecase

		user *controller.UserController

		w *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		input := usecase.SignInInput{
			Name:     "admin",
			Password: "admin",
		}

		output = usecase.SignInOutput{
			ID:       1,
			Name:     "admin",
			Nickname: "administrator",
		}

		mockCtrl = gomock.NewController(GinkgoT())
		mockUserUsecase = mock_usecase.NewMockUserUsecase(mockCtrl)
		mockUserUsecase.EXPECT().SignIn(input).Return(output, nil)

		user = controller.NewUserController(mockUserUsecase)

		w = httptest.NewRecorder()
	})

	Describe("Test SignIn", func() {
		Context("The case that is valid user", func() {
			It("should be status ok", func() {
				expected, _ := json.Marshal(output)

				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("POST", "/user/signin", bytes.NewBufferString("{\"name\":\"admin\", \"password\":\"admin\"}"))

				user.SignIn(c)
				Expect(w.Code).To(Equal(200))
				Expect(w.Body.Bytes()).To(Equal(expected))
			})
		})

		Context("The case that does not match type of name", func() {
			It("should return binding error", func() {
				output := map[string]interface{}{
					"message": "invalid parameter",
				}
				expected, _ := json.Marshal(output)

				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("POST", "/user/signin", bytes.NewBufferString("{\"name\":1, \"password\":\"admin\"}"))

				user.SignIn(c)
				Expect(w.Code).To(Equal(400))
				Expect(w.Body.Bytes()).To(Equal(expected))
			})
		})

		Context("The case that is missing required form", func() {
			It("should return validation error", func() {
				output := map[string]interface{}{
					"message": "invalid parameter",
				}
				expected, _ := json.Marshal(output)

				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("POST", "/user/signin", bytes.NewBufferString("{\"password\":\"admin\"}"))

				user.SignIn(c)
				Expect(w.Code).To(Equal(400))
				Expect(w.Body.Bytes()).To(Equal(expected))
			})
		})
	})
})
