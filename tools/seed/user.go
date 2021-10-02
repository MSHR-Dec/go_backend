package main

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/MSHR-Dec/go_backend/internal/domain/model"
)

type user struct {
	db *gorm.DB
}

func newUser(db *gorm.DB) *user {
	return &user{
			db: db,
	}
}

func (u user) generate() {
	var values []map[string]interface{}

	for i := 0; i < 5; i++ {
		id := strconv.Itoa(i+1)
		v := map[string]interface{}{
			"Name": fmt.Sprintf("fake%s", id),
			"Password": "fake",
			"Nickname": fmt.Sprintf("nickname%s", id),
		}
		values = append(values, v)
	}
	u.db.Clauses(clause.OnConflict{DoNothing: true}).Model(&model.User{}).Create(values)
}
