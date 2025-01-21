package server

import (
	"github.com/onihilist/WebAPI/pkg/api"
	"github.com/onihilist/WebAPI/pkg/databases"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	db := databases.DatabaseConnect()
	databases.DatabaseHealthCheck(db)
	// gin.DisableConsoleColor()

	r := gin.Default()
	//r.SetTrustedProxies(nil)

	r.GET("/ping", api.Ping)
	r.GET("/profile/:name", func(c *gin.Context) {
		api.GetUserProfile(c, db)
	})

	adminAuth := r.Group("/", gin.BasicAuth(GetMiddlewareAdminAuth(db)))
	adminAuth.POST("admin", MiddlewareAdmin)

	return r
}
