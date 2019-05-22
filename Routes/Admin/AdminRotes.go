package Admin

import (
	"BoilerPlateWithAuthInGo/Controllers/Admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	r.GET("/dashboard", Admin.Dashboard)
}
