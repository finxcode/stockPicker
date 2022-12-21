package global

import (
	"go.uber.org/zap"
	"stockPicker/config"
)

type Application struct {
	Logger *zap.Logger
	Config *config.Config
}

var App = new(Application)
