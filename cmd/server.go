package cmd

import (
	"fiber/internal"
	"fmt"
	"os"
)

func startServer() {
	if err := initGlobal(); err != nil {
		return
	}

	app := internal.NewRouter()
	if err := app.Listen(":3000"); err != nil {
		fmt.Println("server run err: ", err)
		os.Exit(ExitServer)
		return
	}

	fmt.Println("server turned off")
}
