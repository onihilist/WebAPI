package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"username": c.Param("name"),
	})
}
