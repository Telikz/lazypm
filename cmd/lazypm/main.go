package main

import (
	"log"
	"os"

	"lazypm/internal/commands"
)

func main() {
	log.SetFlags(0)

	if err := commands.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
