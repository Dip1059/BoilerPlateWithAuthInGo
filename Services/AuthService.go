package Services

import (
	G "BoilerPlateWithAuthInGo/Globals"
	H "BoilerPlateWithAuthInGo/Helpers"
	Mod "BoilerPlateWithAuthInGo/Models"
	R "BoilerPlateWithAuthInGo/Repositories"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)


func SetRememberToken(user Mod.User, c *gin.Context,sc *securecookie.SecureCookie) bool {
	val := user.Email
	encoded, _ := sc.Encode("remember_token", val)

	cookie1 := http.Cookie{
		Name:     "remember_token",
		Value:    encoded,
		MaxAge:   60 * 60 * 24 * 365,
	}

	http.SetCookie(c.Writer, &cookie1)
	user.RememberToken.String = encoded
	user.RememberToken.Valid = true
	if !R.SetRememberToken(user) {
		return false
	}
	return true
}


func SendVerificationEmail(user Mod.User) bool {
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
		G.Msg.Fail = "Verification Email Not Sent, <a href='http://localhost:2000/resend-email-verification'>Click Here To Resend</a>."
		return false
	}
	return true
}


func SendPasswordResetLinkEmail(user Mod.User, ps Mod.PasswordReset) bool {
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
		G.Msg.Fail = "Failed To Send Link, Please Try Again Later."
		return false
	}
	return true
}



