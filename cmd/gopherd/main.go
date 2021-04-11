package main

import (
	"github.com/gopherd/doge/service"

	"github.com/gopherd/demo/cmd/gopherd/server"
)

func main() {
	service.Run(server.New())
}
