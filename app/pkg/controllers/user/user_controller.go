package user

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
	"github.com/onihilist/WebAPI/pkg/databases"
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
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
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
			req := `UPDATE users SET session_id=? WHERE username=?;`
			databases.DoRequest(db, req, encodedSessionID, username) // MAKE CONTROLLER/REPO/SERVICE FOR "sessions"
			c.JSON(http.StatusOK, gin.H{"message": "You are logged in !"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		}

	}
}
