package Routes

import (
	"BoilerPlateWithAuthInGo/Routes/Auth"
	"BoilerPlateWithAuthInGo/Routes/Admin"
	"BoilerPlateWithAuthInGo/Routes/User"
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.LoadHTMLGlob("Views/**/*.html")
	r.Static("/assets", "./")

	Auth.AuthRoutes(r)
	Admin.AdminRoutes(r)
	User.UserRoutes(r)

	r.Run(":2000")
}
