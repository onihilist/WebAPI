package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onihilist/WebAPI/pkg/databases"
	_ "modernc.org/sqlite"
)

type User struct {
	ID             int
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

func CreateUserProfile(c *gin.Context, db *sql.DB) {

}
