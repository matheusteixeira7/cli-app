package main

import (
	"log"
	"os"

	"github.com/matheusteixeira7/cli-app/app"
)

func main() {
	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
