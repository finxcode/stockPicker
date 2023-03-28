package global

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type Application struct {
	Logger *zap.Logger
}

var App = new(Application)
