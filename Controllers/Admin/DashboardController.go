package Admin

import (
	M "BoilerPlateWithAuthInGo/Middlewares"
	G "BoilerPlateWithAuthInGo/Globals"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {

	if user, success := M.IsAuthAdminUser(c, G.Store); success {
		c.HTML(http.StatusOK, "dashboard.html", user)
	}
	return
}
