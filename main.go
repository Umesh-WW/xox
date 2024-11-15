package main

import (
	"github.com/Umesh-WW/xox/api"
)

func main() {
	server := api.NewApiServer(":8080")
	server.Run()
}
