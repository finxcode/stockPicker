package main

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"stockPicker/config"
	"stockPicker/global"
	"stockPicker/internal/db"
	"stockPicker/internal/log"
)

func main() {

	err, c := config.New()
	if err != nil {
		panic(fmt.Sprintf("config initialization failed with error: %s", err.Error()))
	}
	global.App.Config = c
	global.App.Logger = log.New(global.App.Config)
	global.App.Logger.Info("initialization", zap.String("logger", "started"))
	err = db.InitDb(c)
	if err != nil {
		global.App.Logger.Fatal("db connect failed with error", zap.String("fatal", err.Error()))
	}

	defer func() {
		err := global.App.Db.Close()
		if err != nil {
			global.App.Logger.Fatal("db connection close failed with error", zap.String("fatal", err.Error()))
		}
	}()

	//symbols, err := data.FetchStockSymbol()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(len(*symbols))

	for i := 0; i < 10; i++ {
		id := uuid.New()
		fmt.Printf("%s %s\n", id, id.Version().String())
	}

	for i := 0; i < 10; i++ {
		id2, err := uuid.NewRandom()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%s %s\n", id2, id2.Version().String())
	}

}
