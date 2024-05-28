package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed", zap.Error(err))
		return
	}
	DB.SetMaxOpenConns(viper.GetInt("mysql.max_open-conns"))
	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	fmt.Println("connecting to MySql")
	return
}

func Close() {
	_ = DB.Close()
}
