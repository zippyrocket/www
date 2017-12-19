package main

import (
	"log"

	"zippyrocket.com/www/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
