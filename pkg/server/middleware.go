package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMiddlewareAdminAuth() gin.Accounts {
	return gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}
}

func MiddlewareAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}
