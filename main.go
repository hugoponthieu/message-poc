package main

import (
	"message/app"
	"message/cli"
)

func main() {
	config, err := cli.GetConfig()
	if err != nil {
		panic(err)
	}
	server, err := app.InitApp(*config)
	if err != nil {
		panic(err)
	}
	server.Start()
}
