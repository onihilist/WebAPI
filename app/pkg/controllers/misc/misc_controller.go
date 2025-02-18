package misc

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	user "github.com/onihilist/WebAPI/pkg/controllers/user"
	service "github.com/onihilist/WebAPI/pkg/services/misc"
	"github.com/onihilist/WebAPI/pkg/utils"
)

type MiscController struct {
	uc          *user.UserController
	MiscService service.MiscService
}

func NewMiscController(uc *user.UserController, service service.MiscService) MiscController {
	return MiscController{uc, service}
}

func (mc *MiscController) Index(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("session_id")

	user, err := mc.uc.UserService.GetUserBySessionID(sessionID)

	if user.ID != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"SessionID": sessionID,
			"AvatarURL": user.AvatarURL,
			"Username":  user.Username,
		})
	} else if err != nil {
		utils.LogError("%s", err)
	} else {
		utils.LogFatal("Internal server error")
	}
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
