package main

import (
	"fmt"
	"stockPicker/config"
)

func main() {
	err, c := config.New()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Username)

}
