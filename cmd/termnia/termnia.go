package main

import (
	"log"

	"github.com/mateus-luciano/termnia/internal/core"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatal(err)
	}
}
