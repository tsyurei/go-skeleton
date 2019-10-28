package app

import (
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

type Database struct {
}

func InitDatabase() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     viper.GetString("database.host") + ":" + viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	})
}
