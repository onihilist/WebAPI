package server

import (
	"github.com/gin-gonic/gin"
	"github.com/onihilist/WebAPI/pkg/databases"
)

func SetupRouter() *gin.Engine {

	// gin.DisableConsoleColor()

	db := databases.DatabaseConnect()
	databases.DatabaseHealthCheck(db)

	return LoadRoutes(db)

}
