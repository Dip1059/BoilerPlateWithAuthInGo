package Auth

import (
	"BoilerPlateWithAuthInGo/Controllers/Auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	r.GET("/", Auth.Welcome)
	r.GET("/register", Auth.RegisterGet)
	r.POST("/register", Auth.RegisterPost)
	r.GET("resend-email-verification", Auth.ResendEmailVf)
	r.GET("email-verify/:encEmail/:emailVf", Auth.ActivateAccount)
	r.GET("/login", Auth.LoginGet)
	r.POST("/login", Auth.LoginPost)
	r.GET("/forgot-password", Auth.ForgotPassword)
	r.POST("/send-password-reset-link", Auth.SendPasswordResetLink)
	r.GET("/reset-password/:email/:token", Auth.ResetPasswordGet)
	r.POST("/reset-password", Auth.ResetPasswordPost)
	r.GET("/logout", Auth.Logout)
}
