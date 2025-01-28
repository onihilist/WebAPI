package misc

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	miscRepository "github.com/onihilist/WebAPI/pkg/repositories/misc"
	"github.com/onihilist/WebAPI/pkg/utils"
)

type MiscService struct {
	MiscRepository miscRepository.MiscRepository
}

func LoginAdmin(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "Invalid input",
			})
			return
		}

		var storedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", json.Username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  "Invalid username or password",
				})
			} else {
				utils.LogFatal("[MariaDB] - %s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": http.StatusInternalServerError,
					"error":  "Internal server error.",
				})
			}
			return
		}

		if storedPassword != json.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Invalid username or password",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Login successful",
		})
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
	})
}
