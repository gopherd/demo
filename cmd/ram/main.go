package main

import (
	"github.com/gopherd/doge/service"

	"github.com/gopherd/demo/cmd/ram/server"
)

func main() {
	service.Run(server.New())
}
