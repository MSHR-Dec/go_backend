package persistence_test

import (
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/MSHR-Dec/go_backend/internal/domain/model"
	. "github.com/MSHR-Dec/go_backend/internal/interfaces/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ = Describe("Task Repository", func() {
	var (
		mock sqlmock.Sqlmock
		gdb  *gorm.DB
		repo *TaskRepository
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

		repo = NewTaskRepository(gdb)
	})

	Describe("Test ListByUserID", func() {
		Context("Issue the collect SQL", func() {
			It("should return model of User", func() {
				taskModel := []model.Task{
					{
						ID:      1,
						Name:    "admin",
						Summary: "A sample task for admin",
						UserID:  1,
					},
				}

				rows := sqlmock.NewRows([]string{"id", "name", "summary", "user_id"}).AddRow(taskModel[0].ID, taskModel[0].Name, taskModel[0].Summary, taskModel[0].UserID)

				const selected = "SELECT * FROM `tasks` WHERE user_id = ?"

				mock.ExpectQuery(selected).WithArgs(taskModel[0].UserID).WillReturnRows(rows)

				task, err := repo.ListByUserID(1)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(task).Should(Equal(taskModel))
			})
		})

		Context("Issue the invalid SQL", func() {
			It("should return gorm error", func() {
				mock.ExpectQuery("SELECT * FROM `tasks` WHERE user_id = ?").WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{}))
				_, err := repo.ListByUserID(0)
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})
	})
})
