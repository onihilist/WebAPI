package server

import (
	"database/sql"
	"net/http"
	"time"

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

		if userData["userId"] != nil {
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
		}
	})
	r.GET("/profile/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create-user.html", nil)
	})
	r.POST("/profile/create/submit", func(c *gin.Context) {

		phone := c.PostForm("phone")
		var phonePtr *string
		if phone != "" {
			phonePtr = &phone
		}

		user := api.User{
			PermissionID:   3,
			Username:       c.PostForm("username"),
			Password:       c.PostForm("password"),
			Email:          c.PostForm("email"),
			Phone:          phonePtr,
			CreationDate:   time.Now().GoString(),
			LastConnection: time.Now().GoString(),
			LastIP:         "127.0.0.1", // change this for real
		}

		api.CreateUserProfile(db, user)

		c.Redirect(http.StatusSeeOther, "/profile/"+c.PostForm("username"))

	})

	adminAuth := r.Group("/", gin.BasicAuth(GetMiddlewareAdminAuth(db)))
	adminAuth.POST("admin", MiddlewareAdmin)

	return r
}
