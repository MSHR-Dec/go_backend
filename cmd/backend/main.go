// $ go get -u github.com/cosmtrek/air
// $ air -c air.toml
package main

import (
	"github.com/MSHR-Dec/go_backend/internal/infrastructure"
)

func main() {
	env := infrastructure.InitEnv()
	db := infrastructure.InitMySQL(env.MysqlUser, env.MysqlPassword, env.MysqlHost, env.MysqlDatabase)
	cache := infrastructure.InitRedis(env.RedisHost)

	engine := infrastructure.InitGin(db, cache)
	engine.Run(":8080")
}
