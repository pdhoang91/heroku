package main

import (
	"heroku/config"
	"heroku/internal/app"
)

func main() {

	cfg := config.NewConfig()
	app := app.InitializeHTTPServer(cfg)
	app.Run(cfg.HTTPPort)
}
