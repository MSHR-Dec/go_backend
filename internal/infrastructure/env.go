package infrastructure

import (
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Env           string `default:"local"`
	MysqlUser     string `default:"patune" split_words:"true"`
	MysqlPassword string `default:"patune" split_words:"true"`
	MysqlDatabase string `default:"patune" split_words:"true"`
	MysqlHost     string `default:"127.0.0.1:53306" split_words:"true"`
	RedisHost     string `default:"127.0.0.1:6379" split_words:"true"`
}

func InitEnv() Environment {
	var env Environment
	envconfig.Process("", &env)

	return env
}
