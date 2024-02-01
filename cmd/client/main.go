package main

import (
	"os"

	"github.com/mukkak/go-playground/internal/client"
)

func main() {
	err := client.Execute()
	if err != nil {
		os.Exit(1)
	}
}
