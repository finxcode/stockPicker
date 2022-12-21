package main

import (
	"fmt"
	"stockPicker/config"
	"stockPicker/global"
	"stockPicker/internal/log"
)

func main() {

	err, c := config.New()
	if err != nil {
		panic(fmt.Sprintf("config initialization failed with error: %s", err.Error()))
	}
	global.App.Config = c
	global.App.Logger = log.New(global.App.Config)

}
