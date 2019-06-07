package Auth

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	G "BoilerPlateWithAuthInGo/Globals"
	H "BoilerPlateWithAuthInGo/Helpers"
	M "BoilerPlateWithAuthInGo/Middlewares"
	Mod "BoilerPlateWithAuthInGo/Models"
	R "BoilerPlateWithAuthInGo/Repositories"
	S "BoilerPlateWithAuthInGo/Services"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

func Welcome(c *gin.Context) {

	if _, success := M.IsGuest(c, G.Store); success {
		msg := Cfg.GetMessage(c)
		c.HTML(http.StatusOK, "welcome.html", msg)

	}
	return
}

func RegisterGet(c *gin.Context) {

	if _, success := M.IsGuest(c, G.Store); success {
		msg := Cfg.GetMessage(c)
		c.HTML(http.StatusOK, "register.html", msg)

	}
	return
}

func RegisterPost(c *gin.Context) {
	var success bool
	var user Mod.User
	user.FullName = c.PostForm("full_name")
	user.Email = c.PostForm("email")
	_, success = R.ReadWithEmail(user)
	if success {
		Cfg.CreateMessage("fail", "User Already Exists With This Email.", c)
		c.Redirect(http.StatusFound, "/register")
		return
	}
	password := c.PostForm("password")
	confirmPass := c.PostForm("confirm-password")
	if password != confirmPass {
		Cfg.CreateMessage("fail", "Confirm Password Doesn't Match.", c)
		c.Redirect(http.StatusFound, "/register")
		return
	}
	cost := bcrypt.DefaultCost
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = string(hash)
	user.EmailVerification.String = H.RandomString(60)
	user.EmailVerification.Valid = true
	user.RoleID = 2

	user, success = R.Register(user)
	if success {
		if S.SendVerificationEmail(user, c) {
			H.SetCookie("secret", nil, user.Email, "email", 60*60*5, c)
			var link string
			link = "<a href='" + G.AppEnv.Url + "/resend-email-verification'>Click Here To Re-Send Verification Email</a>"
			Cfg.CreateMessage(
				"success", "Successfully Registered. Please Check Your Verification Email. If You Don't Get it "+link+".", c)
		}
		c.Redirect(http.StatusFound, "/register")
	} else {

		Cfg.GetAndSetMessage("fail", "Some Internal Server Error Occurred, Please Try Again.", c)

		c.Redirect(http.StatusFound, "/register")
	}
}

func ResendEmailVf(c *gin.Context) {
	var user Mod.User
	user.Email = H.GetCookie("secret", nil, "email", c)
	user, _ = R.ReadWithEmail(user)
	if user.Email != "" {
		if user.ActiveStatus == 0 {
			if S.SendVerificationEmail(user, c) {
				Cfg.CreateMessage("success", "Email Has Been Sent Successfully.", c)
			}
		} else if user.ActiveStatus > 0 {
			Cfg.CreateMessage("success", "Already Activated.", c)
		} else {
			Cfg.CreateMessage("fail", "Email Sending Failed. Try Again.", c)
		}
	}
	c.Redirect(http.StatusFound, "/login")
}

func ActivateAccount(c *gin.Context) {
	var user Mod.User
	encEmail := c.Param("encEmail")
	emailVf := c.Param("emailVf")
	var err error
	var decoded []byte
	decoded, err = base64.URLEncoding.DecodeString(encEmail)
	if err != nil {
		log.Println("AuthController.go Log1", err.Error())
		c.HTML(http.StatusOK, "404.html", nil)
		return
	}

	user.Email = string(decoded)
	user.EmailVerification.String = emailVf
	var success bool

	user, success = R.ActivateAccount(user)
	if success {
		H.SetCookie("secret", nil, "", "email", -1, c)
		Cfg.CreateMessage("success", "Congratulations, Your Account Is Activated.", c)
		c.Redirect(http.StatusFound, "/login")
	} else {
		c.HTML(http.StatusOK, "404.html", nil)
	}
}

func LoginGet(c *gin.Context) {

	if _, success := M.IsGuest(c, G.Store); success {
		msg := Cfg.GetMessage(c)
		c.HTML(http.StatusOK, "login.html", msg)

	}
	return
}

func LoginPost(c *gin.Context) {
	var user Mod.User

	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	rememberMe, _ := strconv.Atoi(c.PostForm("remember_me"))
	var success bool
	user, success = R.Login(user)
	if success {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			Cfg.CreateMessage("fail", "Wrong Credentials.", c)
			c.Redirect(http.StatusFound, "/login")
		} else {
			if user.ActiveStatus == 1 {
				H.SetCookie("secret", nil, "", "email", -1, c)

				user.RememberToken.String = H.RandomString(60)
				user.RememberToken.Valid = true

				if !R.SetRememberToken(user) {
					Cfg.CreateMessage("fail", "Some Internal Server Error Occurred. Please Try Again.", c)
					c.Redirect(http.StatusFound, "/login")
					return
				}

				session, _ := G.Store.Get(c.Request, "login_token")
				session.Values["userEmail"] = user.Email
				session.Values["remember_token"] = user.RememberToken.String
				session.Options.MaxAge = 60 * 60 * 24 * 5
				session.Save(c.Request, c.Writer)

				if rememberMe == 1 {
					session.Options.MaxAge = 60 * 60 * 24 * 365
					session.Save(c.Request, c.Writer)
				}
				if user.RoleID == 1 {
					c.Redirect(http.StatusFound, "/dashboard")
				} else if user.RoleID == 2 {
					c.Redirect(http.StatusFound, "/home")
				}
			} else if user.ActiveStatus == 2 {

				Cfg.GetAndSetMessage("fail", "You Are Suspended. Contact With The Authority Quickly.", c)

				c.Redirect(http.StatusFound, "/login")
			} else {
				H.SetCookie("secret", nil, user.Email, "email", 60*60*5, c)
				var link string
				link = "<a href='" + G.AppEnv.Url + "/resend-email-verification'>Click Here To Send Verification Email</a>"

				Cfg.GetAndSetMessage("fail", "Please Activate Your Account, "+link+".", c)
				c.Redirect(http.StatusFound, "/login")
			}
		}

	} else {

		Cfg.GetAndSetMessage("fail", "User Not Found.", c)

		c.Redirect(http.StatusFound, "/login")
	}
}

func ForgotPassword(c *gin.Context) {

	if _, success := M.IsGuest(c, G.Store); success {
		msg := Cfg.GetMessage(c)
		c.HTML(http.StatusOK, "forgot-password.html", msg)

	}
}

func SendPasswordResetLink(c *gin.Context) {
	var user Mod.User
	var ps Mod.PasswordReset
	var success bool
	user.Email = c.PostForm("email")
	user, success = R.ReadWithEmail(user)
	if !success {
		Cfg.CreateMessage("fail", "User Not Found With This Email.", c)
		c.Redirect(http.StatusFound, "/forgot-password")
		return
	}
	if user.ActiveStatus == 0 {
		H.SetCookie("secret", nil, user.Email, "email", 60*60*5, c)
		var link string
		link = "<a href='" + G.AppEnv.Url + "/resend-email-verification'>Click Here To Send Verification Email</a>"

		Cfg.GetAndSetMessage("fail", "Please Activate Your Account, "+link+".", c)

		c.Redirect(http.StatusFound, "/login")
		return
	}
	if user.ActiveStatus == 2 {
		Cfg.CreateMessage("fail", "User Is Suspended.", c)
		c.Redirect(http.StatusFound, "/login")
		return
	}
	ps.Email = user.Email
	ps.Token.String = H.RandomString(60)
	ps.Token.Valid = true
	if !R.SendPasswordResetLink(ps) {
		return
	}
	if S.SendPasswordResetLinkEmail(user, ps, c) {
		Cfg.CreateMessage("success", "Reset Password Link Sent Successfully. Check Your Email.", c)
	}
	c.Redirect(http.StatusFound, "/login")
}

func ResetPasswordGet(c *gin.Context) {
	if _, success := M.IsGuest(c, G.Store); !success {
		return
	}
	encEmail := c.Param("email")
	var err error
	var decoded []byte
	decoded, err = base64.URLEncoding.DecodeString(encEmail)
	if err != nil {
		log.Println("AuthController.go Log2", err.Error())
		c.HTML(http.StatusOK, "404.html", nil)
		return
	}
	var ps Mod.PasswordReset
	ps.Email = string(decoded)
	ps.Token.String = c.Param("token")
	ps.Token.Valid = true
	if !R.ResetPasswordGet(ps) {
		c.HTML(http.StatusOK, "404.html", nil)
		return
	}
	msg := Cfg.GetMessage(c)
	c.HTML(http.StatusOK, "reset-password.html", map[string]interface{}{"Msg": msg, "PS": ps})

}

func ResetPasswordPost(c *gin.Context) {
	var user Mod.User
	var ps Mod.PasswordReset
	password := c.PostForm("password")
	confirmPass := c.PostForm("confirm-password")
	ps.Email = c.PostForm("email")
	ps.Token.String = c.PostForm("token")
	ps.Token.Valid = true
	if password != confirmPass {
		Cfg.CreateMessage("fail", "Confirm Password Doesn't Match.", c)
		encEmail := base64.URLEncoding.EncodeToString([]byte(ps.Email))
		c.Redirect(http.StatusFound, "/reset-password/"+encEmail+"/"+ps.Token.String)
		return
	}
	cost := bcrypt.DefaultCost
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = string(hash)
	user.Email = ps.Email
	if !R.ResetPasswordPost(user, ps) {
		Cfg.CreateMessage("fail", "Some Internal Server Error Occurred, Please Try Again Later.", c)
		encEmail := base64.URLEncoding.EncodeToString([]byte(ps.Email))
		c.Redirect(http.StatusFound, "/reset-password/"+encEmail+"/"+ps.Token.String)
		return
	}
	Cfg.CreateMessage("success", "Your Password Is Reset Successfully.", c)
	c.Redirect(http.StatusFound, "/login")
}

func Logout(c *gin.Context) {
	var user Mod.User

	session, _ := G.Store.Get(c.Request, "login_token")
	user.Email = session.Values["userEmail"].(string)
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	R.Logout(user)
	c.Redirect(http.StatusFound, "/login")
}
