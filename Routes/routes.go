package Routes

import (
	Cont "BoilerPlateWithAuthInGo/Controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.LoadHTMLGlob("Views/**/*.html")
	r.GET("/", Cont.Welcome)
	r.GET("/register", Cont.RegisterGet)
	r.POST("/register", Cont.RegisterPost)
	r.GET("resend-email-verification", Cont.ResendEmailVf)
	r.GET("email-verify/:encEmail/:emailVf", Cont.ActivateAccount)
	r.GET("/login", Cont.LoginGet)
	r.POST("/login", Cont.LoginPost)
	r.GET("/forgot-password", Cont.ForgotPassword)
	r.POST("/send-password-reset-link", Cont.SendPasswordResetLink)
	r.GET("/reset-password/:email/:token", Cont.ResetPasswordGet)
	r.POST("/reset-password", Cont.ResetPasswordPost)
	r.GET("/home", Cont.Home)
	r.GET("/dashboard", Cont.Dashboard)
	r.GET("/logout", Cont.Logout)

	r.Run(":2000")
}
