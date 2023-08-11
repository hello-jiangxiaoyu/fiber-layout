package main

import (
	"fiber/internal"
)

func main() {
	if err := internal.InitConfig(); err != nil {
		return
	}
	if err := internal.InitLogger(); err != nil {
		return
	}
	app := internal.NewRouter()

	if err := app.Listen(":3000"); err != nil {
		return
	}
}
