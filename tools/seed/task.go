package main

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
)

type task struct {
	db *gorm.DB
}

func newTask(db *gorm.DB) *task {
	return &task{
			db: db,
	}
}

func (t task) generate() {
	userIDs := t.getUserIDs()

	for _, userID := range userIDs {
		for i := 0; i < 20; i++ {
			id := strconv.Itoa(i+1)
			task := model.Task{
				Name: fmt.Sprintf("fake%s", id),
				Summary: fmt.Sprintf("A fake task %s", id),
				UserID: userID,
			}
			t.db.FirstOrCreate(&model.Task{}, task)
		}
	}
}

func (t task) getUserIDs() []int {
	var userIDs []int

	rows, _ := t.db.Table("users").Select("id").Where("password = ?", "fake").Rows()
	for rows.Next() {
		t.db.ScanRows(rows, &userIDs)
	}

	return userIDs
}
