package main

import (
	"log"
	"os"

	"github.com/wanomir/hh-reporter/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(a.Run())
}
