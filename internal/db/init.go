package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"stockPicker/config"
	"stockPicker/global"
)

func InitDb(c *config.Config) error {
	user := c.DB.Username
	password := c.DB.Password
	host := c.DB.Host
	port := c.DB.Port
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/database", user, password, host, port)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		global.App.Logger.Error("connect server failed, err:%v\n", zap.String("connect db error", err.Error()))
		return err
	}
	global.App.Db = db
	return nil
}
