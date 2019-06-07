package Services

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	H "BoilerPlateWithAuthInGo/Helpers"
	Mod "BoilerPlateWithAuthInGo/Models"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
)


func SendVerificationEmail(user Mod.User, c *gin.Context) bool {
	encEmail := base64.URLEncoding.EncodeToString([]byte(user.Email))
	htmlString, err := H.ParseTemplate("Views/Email/email-verify.html", map[string]interface{}{
		"EncEmail":encEmail, "User":user})
	if err != nil {
		log.Println("AuthService.go Log1", err.Error())
		return false
	}

	From := "Gophers <gopher@mail.com>"
	To := user.Email
	Subject := "Account Verification Email"
	HtmlString := htmlString
	if !SendEmail(From, To, Subject, HtmlString) {
		Cfg.CreateMessage("fail","Verification Email Not Sent, Try Again Later.",c)
		return false
	}
	return true
}


func SendPasswordResetLinkEmail(user Mod.User, ps Mod.PasswordReset,c *gin.Context) bool {
	encEmail := base64.URLEncoding.EncodeToString([]byte(user.Email))

	htmlString, err := H.ParseTemplate("Views/Email/reset-password-email.html", map[string]interface{}{
		"EncEmail":encEmail, "User":user, "PS":ps})
	if err != nil {
		log.Println("AuthService.go Log2", err.Error())
		return false
	}

	From := "Gophers <gopher@mail.com>"
	To := user.Email
	Subject := "Reset Password Link"
	HtmlString := htmlString
	if !SendEmail(From, To, Subject, HtmlString) {
		Cfg.CreateMessage("fail","Failed To Send Link, Please Try Again Later.",c)
		return false
	}
	return true
}



