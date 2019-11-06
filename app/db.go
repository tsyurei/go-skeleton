package app

import (
	"fmt"

	"github.com/go-pg/pg/v9"
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

func GetDbURL() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.database"))
}
