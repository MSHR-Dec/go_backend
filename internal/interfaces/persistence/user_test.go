package persistence_test

import (
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	. "github.com/MSHR-Dec/go_backend/internal/interfaces/persistence"
)

var _ = Describe("User Repository", func() {
	var (
		mock sqlmock.Sqlmock
		gdb *gorm.DB
		repo *UserRepository
	)

	BeforeEach(func() {
		var (
			db  *sql.DB
			err error
		)

		db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err = gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())

		repo = NewUserRepository(gdb)
	})

	Describe("Test GetByName", func() {
		Context("Issue the collect SQL", func() {
			It("should return model of User", func() {
				userModel := &model.User{
					ID:       1,
					Name:     "admin",
					Password: "admin",
				}

				rows := sqlmock.NewRows([]string{"id", "name", "password"}).AddRow(userModel.ID, userModel.Name, userModel.Password)

				const selected = "SELECT * FROM `users` WHERE name = ? ORDER BY `users`.`id` LIMIT 1"

				mock.ExpectQuery(selected).WithArgs(userModel.Name).WillReturnRows(rows)

				user, err := repo.GetByName(userModel.Name)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(user).Should(Equal(*userModel))
			})
		})

		Context("Issue the invalid SQL", func() {
			It("should return gorm error", func() {
				mock.ExpectQuery("SELECT * FROM `users` WHERE name = ? ORDER BY `users`.`id` LIMIT 1").WithArgs("dummy").WillReturnRows(sqlmock.NewRows([]string{}))
				_, err := repo.GetByName("dummy")
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})
	})
})
