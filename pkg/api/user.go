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

func GetUserProfile(c *gin.Context, db *sql.DB) {
	username := c.Param("name")

	query := "SELECT id, username, password, email, phone, creationDate, lastConnection, lastIP FROM users WHERE username = ?"

	row := databases.DoRequestRow(db, query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.CreationDate, &user.LastConnection, &user.LastIP)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User  not found"})
			return
		}
		log.Fatal(err)
	}

	if user.Phone == nil {
		fmt.Println("Phone number is not available")
	}

	c.JSON(http.StatusOK, user)

}
