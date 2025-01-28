package entities

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onihilist/WebAPI/pkg/databases"
	"github.com/onihilist/WebAPI/pkg/utils"
)

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	PermissionID   uint      `json:"permissionId" gorm:"not null"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Password       string    `json:"-" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Phone          *string   `json:"phone,omitempty" gorm:"unique"`
	CreationDate   time.Time `json:"creationDate" gorm:"not null"`
	LastConnection time.Time `json:"lastConnection" gorm:"not null"`
	LastIP         string    `json:"lastIP" gorm:"not null"`
	SessionID      *string   `json:"session_id,omitempty" gorm:"size:512"`
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
