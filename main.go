package main

import (
	"log"
	"network-search/app"
	"os"
)

func main() {
	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
