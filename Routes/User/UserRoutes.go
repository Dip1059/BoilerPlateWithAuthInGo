package User

import (
	"BoilerPlateWithAuthInGo/Controllers/User"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/home", User.Home)
}
