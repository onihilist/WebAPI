package main

import (
	"github.com/onihilist/WebAPI/pkg/server"
)

func main() {
	r := server.SetupRouter()
	r.StaticFile("/favicon.ico", "./public/img/favicon.ico")
	r.Run(":8080")
}
