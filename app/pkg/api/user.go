package api

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func DeleteUserProfile(db *sql.DB, username string) {
	req := `DELETE FROM users WHERE username = ?;`
	databases.DoRequest(db, req, username)
}

func LoginUser(c *gin.Context, db *sql.DB, username string, password string) {
	req := `SELECT password FROM users WHERE username=?`
	row := databases.DoRequestRow(db, req, username)

	hashPass := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(hashPass[:])

	var pass string
	err := row.Scan(&pass)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		}
		utils.LogError("[/login/check/%s] - %s", username, err.Error())
	} else {
		if pass == hashString {
			session := sessions.Default(c)
			uniqueSessionID := uuid.New().String()
			encodedSessionID := base64.StdEncoding.EncodeToString([]byte(uniqueSessionID))
			session.Set("session_id", encodedSessionID)
			session.Save()
			req := `UPDATE users SET session_id=? WHERE username=?;`
			databases.DoRequest(db, req, encodedSessionID, username)
			c.JSON(http.StatusOK, gin.H{"message": "You are logged in !"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		}
	}
}

func DeleteSessionCookie(c *gin.Context, db *sql.DB, session interface{}) {
	c.SetCookie("gin_session", "", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie gin_session a été supprimé")
	databases.DoRequest(db, `UPDATE users SET session_id=NULL WHERE session_id=?`, session)
}
