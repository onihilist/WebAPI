package user

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/onihilist/WebAPI/pkg/entities"
	service "github.com/onihilist/WebAPI/pkg/services/user"
	"github.com/onihilist/WebAPI/pkg/utils"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return UserController{}
}

// CreateUser  handles the creation of a new user.
func (uc *UserController) CreateUser(c *gin.Context) {

	phone := c.PostForm("phone")
	var phonePtr *string
	if phone != "" {
		phonePtr = &phone
	} else {
		phonePtr = nil
	}

	user := entities.User{
		ID:             nil,
		PermissionID:   3,
		Username:       c.PostForm("username"),
		Password:       c.PostForm("password"),
		Email:          c.PostForm("email"),
		Phone:          phonePtr,
		CreationDate:   time.Now(),
		LastConnection: time.Now(),
		LastIP:         "127.0.0.1", // change this for real
		SessionID:      nil,
	}

	if err := uc.UserService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, user)
		c.Redirect(http.StatusSeeOther, "/profile/"+c.PostForm("username"))
	}

}

// GetUser Profile handles the request to retrieve a user profile.
func (uc *UserController) GetUser(c *gin.Context) {
	username := c.Param("name")

	user, err := uc.UserService.GetUser(username)
	if err != nil {
		utils.LogError("[/profile/%s] - %s", username, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if user == nil {
		utils.LogWarning("[/profile/%s] - User not found", username)
		c.JSON(http.StatusNotFound, gin.H{"error": "User  not found"})
		return
	}

	if user.Phone == nil {
		utils.LogInfo("[/profile/%s] - Phone number is not available for this user", username)
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"userId":         user.ID,
		"username":       user.Username,
		"password":       user.Password, // Consider omitting this
		"email":          user.Email,
		"phone":          user.Phone,
		"creationDate":   user.CreationDate,
		"lastConnection": user.LastConnection,
		"lastIP":         user.LastIP,
	})
}

func (uc *UserController) GetUserBySessionID(c *gin.Context) {
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

	user, permission, err := uc.UserService.GetUserBySessionID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User  not found"})
		return
	}

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

func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("username")

	if err := uc.UserService.DeleteUser(username); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User  not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil) // 204 No Content
}

func (uc *UserController) LoginUser(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := uc.UserService.GetUser(username)

	if err != nil {

		utils.LogError("[/login/check/%s] - %s", username, err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})

	} else {

		hashPass := md5.Sum([]byte(password))
		hashString := hex.EncodeToString(hashPass[:])

		if user.Password == hashString {
			session := sessions.Default(c)
			uniqueSessionID := uuid.New().String()
			encodedSessionID := base64.StdEncoding.EncodeToString([]byte(uniqueSessionID))
			session.Set("session_id", encodedSessionID)
			session.Save()
			uc.UserService.UpdateSessionCookie(encodedSessionID) // MAKE CONTROLLER/REPO/SERVICE FOR "sessions"
			c.JSON(http.StatusOK, gin.H{"message": "You are logged in !"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		}

	}
}

func (uc *UserController) LoginAdmin(c *gin.Context) {
	var json struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid input",
		})
		return
	}

	isValid, err := uc.UserService.LoginAdmin(json.Username, json.Password)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Invalid username or password",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  "Internal server error.",
			})
		}
		return
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "Invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login successful",
	})
}

func (uc *UserController) Disconnect(c *gin.Context) {

	session := sessions.Default(c)

	c.SetCookie("gin_session", "", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie gin_session a été supprimé")

	sessionID := session.Get("session_id")

	if sessionID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No session found"})
		return
	}

	_, err := uc.UserService.UpdateSessionCookie(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to disconnect"})
		return
	}

	// Optionally, you can clear the session cookie
	session.Delete("session_id")
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Successfully disconnected"})
}
