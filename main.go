package main

import (
	"github.com/shenglol/golang/app"
	"github.com/shenglol/golang/config"
)

func main() {
	config := config.GetConfig()
	app := app.NewApp(config)
	app.Run()
}
