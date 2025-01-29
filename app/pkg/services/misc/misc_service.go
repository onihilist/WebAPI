package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	miscRepository "github.com/onihilist/WebAPI/pkg/repositories/misc"
)

type MiscService struct {
	MiscRepository miscRepository.MiscRepository
}

func NewMiscService(repo miscRepository.MiscRepository) MiscService {
	return MiscService{repo}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
	})
}
