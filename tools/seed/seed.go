package main

import (
	"github.com/MSHR-Dec/go_backend/internal/infrastructure"
)

func main() {
	env := infrastructure.InitEnv()
	db := infrastructure.InitMySQL(env.MysqlUser, env.MysqlPassword, env.MysqlHost, env.MysqlDatabase)

	user := newUser(db)
	user.generate()

	task := newTask(db)
	task.generate()
}
