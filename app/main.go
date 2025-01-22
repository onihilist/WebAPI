package main

import (
	"github.com/onihilist/WebAPI/pkg/server"
)

func main() {
	r := server.SetupRouter()
	r.Run(":8080")
}
