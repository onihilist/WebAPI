package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "github.com/onihilist/WebAPI/pkg/services/misc"
)

type MiscController struct {
	MiscService service.MiscService
}

func NewMiscController(service service.MiscService) MiscController {
	return MiscController{service}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
	})
}

func FormCreateUser(c *gin.Context) {
	c.HTML(http.StatusOK, "create-user.html", nil)
}
