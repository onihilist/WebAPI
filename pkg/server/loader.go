package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onihilist/WebAPI/pkg/api"
)

func LoadRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//r.SetTrustedProxies(nil)

	r.GET("/")
	r.GET("/ping", api.Ping)
	r.GET("/profile/:name", func(c *gin.Context) {
		userData := api.GetUserProfile(c, db)
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"UserID":         userData["userId"],
			"Username":       userData["username"],
			"Password":       userData["password"],
			"Email":          userData["email"],
			"Phone":          userData["phone"],
			"CreationDate":   userData["creationDate"],
			"LastConnection": userData["lastConnection"],
			"LastIP":         userData["lastIP"],
		})
	})

	adminAuth := r.Group("/", gin.BasicAuth(GetMiddlewareAdminAuth(db)))
	adminAuth.POST("admin", MiddlewareAdmin)

	return r
}
