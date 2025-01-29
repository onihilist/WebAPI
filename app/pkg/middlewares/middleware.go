package middlewares

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onihilist/WebAPI/pkg/entities"
	"github.com/onihilist/WebAPI/pkg/utils"
	_ "modernc.org/sqlite"
)

var dataAuth = make(map[string]string)

func GetMiddlewareAdminAuth(db *sql.DB) gin.Accounts {
	rows, err := db.Query("SELECT username, password FROM users WHERE permission_id=1")
	if err != nil {
		utils.LogError("[MariaDB] - %s", err.Error())
		return gin.Accounts{}
	}
	defer rows.Close()

	accounts := gin.Accounts{}
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			utils.LogError("[MariaDB] - %s", err.Error())
			continue
		}
		accounts[user.Username] = user.Password
	}

	return accounts
}

func MiddlewareAdmin(accounts gin.Accounts) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get(gin.AuthUserKey)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		username, ok := user.(string)
		if !ok || accounts[username] == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			c.Abort()
			return
		}

		dataAuth[username] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
		c.Next()
	}
}
