package global

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"stockPicker/config"
)

type Application struct {
	Logger *zap.Logger
	Config *config.Config
	Db     *sqlx.DB
}

var App = new(Application)
