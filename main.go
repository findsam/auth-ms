package main

import (
	"log"

	"github.com/findsam/auth-micro/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
