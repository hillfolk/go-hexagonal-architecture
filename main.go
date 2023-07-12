package main

import "github.com/hillfolk/go-hexagonal-architecture/infra/web"

func main() {
	server := web.NewHttpServer()
	if server.Run() != nil {
		panic("Error running server")
	}
}
