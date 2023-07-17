package main

import (
	database "github.com/hillfolk/go-hexagonal-architecture/infra/db"
	"github.com/hillfolk/go-hexagonal-architecture/infra/web"
)

func main() {
	db := database.Connection()
	server := web.NewHttpServer()
	if server.Run() != nil {
		panic("Error running server")
	}
}
