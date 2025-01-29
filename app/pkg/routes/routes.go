package routes

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/onihilist/WebAPI/pkg/controllers/misc"
	"github.com/onihilist/WebAPI/pkg/middlewares"
)

func LoadRoutes(app *App) *gin.Engine {

	r := gin.Default()
	store := sessions.NewCookieStore([]byte("store_session"))

	r.LoadHTMLGlob("templates/*")
	r.Use(sessions.Sessions("gin_session", store))
	//r.SetTrustedProxies(nil)

	// MISC ROUTES
	r.GET("/")
	r.GET("/ping", misc.Ping)

	// LOGIN ROUTES
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login/check", func(c *gin.Context) {
		app.UserController.LoginUser(c)
	})

	// PROFILE ROUTES
	r.GET("/profile/:name", app.UserController.GetUser)
	r.GET("/profile/create", misc.FormCreateUser)
	r.POST("/profile/create/submit", app.UserController.CreateUser)

	r.GET("/profile/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile-settings.html", nil)
	})

	// ADMIN ROUTES
	adminAccounts := middlewares.GetMiddlewareAdminAuth(app.DB)
	adminAuth := r.Group("/", gin.BasicAuth(adminAccounts))
	adminAuth.POST("/admin/login", middlewares.MiddlewareAdmin(adminAccounts))
	adminAuth.POST("/admin/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Admin Dashboard!"})
	})

	r.GET("/user", app.UserController.GetUserBySessionID)

	r.GET("/disconnect", app.UserController.Disconnect)

	return r
}
