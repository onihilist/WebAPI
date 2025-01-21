package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var VALID_TOKENS = map[string]bool{
	"Bearer 6ddebcd7954bfb7baf0f694c1bb7d243f6c8077e509e1520ebf28b1572ce7be0": true,
	// Add some valid tokens here
}

func CheckBearerToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !VALID_TOKENS[authHeader] {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Bearer OK",
	})
	c.Next()
}

func HandlerMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, authorized user!"))
}
