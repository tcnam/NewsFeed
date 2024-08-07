package main

import "github.com/tcnam/golangbe/pkg/handlers"

func main() {
	const (
		port string = ":8000"
	)
	server := handlers.NewAPIServer(port)
	server.Run()
}
