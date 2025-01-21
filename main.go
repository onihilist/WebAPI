package main

import (
	"net/http"

	"github.com/onihilist/WebAPI/pkg/api"
	"github.com/onihilist/WebAPI/pkg/databases"

	"github.com/gin-gonic/gin"
)

// For const like http.StatusOK

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	databases.KusabaConnect()
	// gin.DisableConsoleColor()

	r := gin.Default()
	//r.SetTrustedProxies(nil)

	r.GET("/ping", api.Ping)
	r.GET("/profile/:name", api.GetUserProfile)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
