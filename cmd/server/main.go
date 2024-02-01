package main

import (
	"log"

	"github.com/mukkak/go-playground/internal/server"
)

func main() {
	log.Fatal(server.StartServer())
}
