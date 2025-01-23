package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onihilist/WebAPI/pkg/api"
	"github.com/onihilist/WebAPI/pkg/utils"
	_ "modernc.org/sqlite"
)

var dataAuth = make(map[string]string)

func GetMiddlewareAdminAuth(db *sql.DB) gin.Accounts {
	rows, err := db.Query("SELECT username, password FROM users WHERE permissionId=1")
	if err != nil {
		utils.LogError("[MariaDB] - %s", err.Error())
		return gin.Accounts{}
	}
	defer rows.Close()

	accounts := gin.Accounts{}
	for rows.Next() {
		var user api.User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			utils.LogError("[MariaDB] - %s", err.Error())
			continue
		}
		accounts[user.Username] = user.Password
	}

	return accounts
}

func MiddlewareAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		dataAuth[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}
