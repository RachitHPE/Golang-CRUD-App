package main

import (
	server "crudApplication/internal/pkg/server"
)

func main() {
	start := server.New("Books")
	start.ConfigureAndStart()
}
