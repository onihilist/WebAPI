package routes

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/onihilist/WebAPI/pkg/databases"
	"github.com/onihilist/WebAPI/pkg/utils"
)

func LoadRoutes(app *App) *gin.Engine {

	r := gin.Default()
	store := sessions.NewCookieStore([]byte("store_session"))

	r.LoadHTMLGlob("templates/*")
	r.Use(sessions.Sessions("gin_session", store))
	//r.SetTrustedProxies(nil)

	r.GET("/")
	r.GET("/ping", app.MiscController.Ping)
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login/check", func(c *gin.Context) {
		app.UserController.LoginUser(c)
	})
	r.GET("/profile/:name", app.UserController.GetUser)
	r.GET("/profile/create")
	r.POST("/profile/create/submit", app.UserController.CreateUser)

	r.GET("/profile/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile-settings.html", nil)
	})

	adminAuth := r.Group("/", gin.BasicAuth(GetMiddlewareAdminAuth(db)))
	adminAuth.POST("/admin/login", MiddlewareAdmin)
	adminAuth.POST("/admin/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Admin Dashboard!"})
	})

	r.GET("/user", func(c *gin.Context) {

		type Response struct {
			Status   int     `json:"status"`
			Role     string  `json:"role"`
			Username string  `json:"username"`
			Email    string  `json:"email"`
			Phone    *string `json:"phone"`
		}

		session := sessions.Default(c)
		userID := session.Get("session_id")

		if userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User  not logged in"})
			return
		}

		var user api.User
		var permission string
		err := databases.DoRequestRow(db, "SELECT username, email, phone, permissionId FROM users WHERE session_id = ?", userID).Scan(&user.Username, &user.Email, &user.Phone, &user.PermissionID)
		errPerm := databases.DoRequestRow(db, "SELECT permission FROM permissions WHERE id = ?", user.PermissionID).Scan(&permission)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User  not found"})
			return
		} else if errPerm != nil {
			utils.LogWarning("Permission unknown : %s", permission)
			c.JSON(http.StatusNotFound, gin.H{"message": "The permission of the user is unknown"})
		} else {
			var phonePtr *string
			if user.Phone != nil && *user.Phone != "" {
				phonePtr = user.Phone
			}
			c.JSON(http.StatusOK, Response{
				Status:   http.StatusOK,
				Role:     permission,
				Username: user.Username,
				Email:    user.Email,
				Phone:    phonePtr,
			})
		}
	})

	r.GET("/disconnect", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("session_id")
		api.DeleteSessionCookie(c, db, sessionID)
	})

	return r
}
