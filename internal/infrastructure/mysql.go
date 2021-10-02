package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(user string, password string, host string, db string) *gorm.DB {
	endpoint := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		db)

	conn, err := gorm.Open(mysql.Open(endpoint), &gorm.Config{})
	if err != nil {
		panic("Fail to connect Database.")
	}

	return conn
}
