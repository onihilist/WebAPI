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

	// LOAD HTML TEMPLATES
	r.LoadHTMLFiles("templates/login.html",
		"templates/create-user.html",
		"templates/profile-settings.html",
		"templates/profile.html",
		"templates/partials/navbar.html",
		"templates/partials/header.html")

	r.Static("/uploads", "./uploads")
	r.Static("/public", "./public")
	r.Use(sessions.Sessions("gin_session", store))
	//r.SetTrustedProxies(nil)

	// MISC ROUTES
	r.GET("/")
	r.GET("/ping", misc.Ping)

	// LOGIN ROUTES
	r.GET("/login", app.UserController.LoginPage)
	r.POST("/login/check", func(c *gin.Context) {
		app.UserController.LoginUser(c)
	})

	// PROFILE ROUTES
	r.GET("/profile/create", misc.FormCreateUser)
	r.GET("/profile/delete" /* create the controller */)
	r.GET("/profile/:name", app.UserController.GetUser)
	r.GET("/profile/settings", app.UserController.UserSettings)
	r.POST("/profile/create/submit", app.UserController.CreateUser)
	r.POST("/profile/upload/avatar", app.UserController.UploadAvatar)

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
