package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise"
)

func main() {
	path := flag.String("config", "", "Appliation config path")
	flag.Parse()

	app := task.NewApplication()

	fmt.Println(*path)

	if err := app.LoadConfig(*path); err != nil {
		log.Fatal(err)
	}

	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
