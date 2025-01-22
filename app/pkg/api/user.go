package api

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/onihilist/WebAPI/pkg/databases"
)

type User struct {
	ID             *int
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
			c.JSON(http.StatusNotFound, gin.H{"error": "User  not found"})
			return nil
		}
		log.Fatal(err)
	}

	if user.Phone == nil {
		fmt.Println("Phone number is not available")
	}

	return gin.H{
		"status":         http.StatusOK,
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
		req := `INSERT INTO users (username, password, email, phone, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?);`
		databases.DoRequest(
			db,
			req,
			user.Username,
			hashString,
			user.Email,
			user.Phone,
			user.CreationDate,
			user.LastConnection,
			user.LastIP,
		)
	} else {
		req := `INSERT INTO users (username, password, email, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?);`
		databases.DoRequest(
			db,
			req,
			user.Username,
			hashString,
			user.Email,
			user.CreationDate,
			user.LastConnection,
			user.LastIP,
		)
	}

}
