package User

import (
	M "BoilerPlateWithAuthInGo/Middlewares"
	G "BoilerPlateWithAuthInGo/Globals"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {

	if user, success := M.IsAuthUser(c, G.Store); success {
		c.HTML(http.StatusOK, "home.html", user)
	}
	return
}
