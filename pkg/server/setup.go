package server

import (
	"github.com/onihilist/WebAPI/pkg/api"
	"github.com/onihilist/WebAPI/pkg/databases"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {

	databases.KusabaConnect()
	// gin.DisableConsoleColor()

	r := gin.Default()
	//r.SetTrustedProxies(nil)

	r.GET("/ping", api.Ping)
	r.GET("/profile/:name", api.GetUserProfile)

	adminAuth := r.Group("/", gin.BasicAuth(GetMiddlewareAdminAuth()))
	adminAuth.POST("admin", MiddlewareAdmin)

	return r
}
