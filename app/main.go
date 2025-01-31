package main

import (
	"github.com/onihilist/WebAPI/pkg/routes"
)

func main() {
	r := routes.SetupRouter()
	r.StaticFile("/favicon.ico", "./public/img/favicon.ico")
	r.Run(":8080")
}
