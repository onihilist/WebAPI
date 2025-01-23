package api

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/onihilist/WebAPI/pkg/databases"
	"github.com/onihilist/WebAPI/pkg/utils"
)

type User struct {
	ID             *int
	PermissionID   int
	Username       string
	Password       string
	Email          string
	Phone          *string
	CreationDate   string
	LastConnection string
	LastIP         string
}

func GetUserProfile(c *gin.Context, db *sql.DB) gin.H {
	username := c.Param("name")

	query := "SELECT id, username, password, email, phone, creationDate, lastConnection, lastIP FROM users WHERE username = ?"

	row := databases.DoRequestRow(db, query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.CreationDate, &user.LastConnection, &user.LastIP)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.LogWarning("[/profile/%s] - User not found", username)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return nil
		}
		utils.LogError("[/profile/%s] - %s", username, err.Error())
	}

	if user.Phone == nil {
		utils.LogInfo("[/profile/%s] - Phone number is not available for this user", username)
	}

	return gin.H{
		"userId":         user.ID,
		"username":       user.Username,
		"password":       user.Password,
		"email":          user.Email,
		"phone":          user.Phone,
		"creationDate":   user.CreationDate,
		"lastConnection": user.LastConnection,
		"lastIP":         user.LastIP,
	}

}

func CreateUserProfile(db *sql.DB, user User) {

	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])

	if user.Phone != nil {
		req := `INSERT INTO users (permissionId, username, password, email, phone, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
		databases.DoRequest(
			db,
			req,
			user.PermissionID,
			user.Username,
			hashString,
			user.Email,
			user.Phone,
			time.Now(),
			time.Now(),
			user.LastIP,
		)
	} else {
		req := `INSERT INTO users (permissionId, username, password, email, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?);`
		databases.DoRequest(
			db,
			req,
			user.PermissionID,
			user.Username,
			hashString,
			user.Email,
			time.Now(),
			time.Now(),
			user.LastIP,
		)
	}

}
