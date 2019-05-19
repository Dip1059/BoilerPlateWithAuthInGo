package Controllers

import (
	G "BoilerPlateWithAuthInGo/Globals"
	H "BoilerPlateWithAuthInGo/Helpers"
	M "BoilerPlateWithAuthInGo/Middlewares"
	"BoilerPlateWithAuthInGo/Models"
	R "BoilerPlateWithAuthInGo/Repositories"
	S "BoilerPlateWithAuthInGo/Services"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


var (
	sc    = securecookie.New([]byte("secret"), nil)
	store = sessions.NewCookieStore([]byte("secret"))
)


func Welcome(c *gin.Context) {
	if M.IsGuest(c, store, sc) {
		c.HTML(http.StatusOK, "welcome.html", nil)
	}
	return
}


func RegisterGet(c *gin.Context) {
	if M.IsGuest(c, store, sc) {
		c.HTML(http.StatusOK, "register.html", G.Msg)
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
	return
}


func RegisterPost(c *gin.Context) {
	var success bool
	G.User.FullName = c.PostForm("full_name")
	G.User.Email = c.PostForm("email")
	_, success = R.ReadWithEmail(G.User)
	if success {
		G.Msg.Fail = "User Already Exists With This Email."
		c.Redirect(http.StatusFound, "/register")
		return
	}
	password := c.PostForm("password")
	confirmPass := c.PostForm("confirm-password")
	if password != confirmPass {
		G.Msg.Fail = "Confirm Password Doesn't Match."
		c.Redirect(http.StatusFound, "/register")
		return
	}
	cost := bcrypt.DefaultCost
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	G.User.Password = string(hash)
	G.User.EmailVerification.String = H.RandomString(60)
	G.User.EmailVerification.Valid = true
	G.User.RoleID = 2

	G.User, success = R.Register(G.User)
	if success {
		if S.SendVerificationEmail() {
			var link template.HTML
			link = "<a href='http://localhost:2000/resend-email-verification'>Click Here To Resend</a>"
			G.Msg.Success = "Successfully Registered. Please Check Your Verification Email. If You Don't Get it " + link + "."
		}
		c.Redirect(http.StatusFound, "/register")
	} else {
		if G.Msg.Fail == "" {
			G.Msg.Fail = "Some Internal Server Error Occurred, Please Try Again."
		}
		c.Redirect(http.StatusFound, "/register")
	}
}


func ResendEmailVf(c *gin.Context) {
	if G.User.Email != "" {
		if G.User.ActiveStatus == 0 {
			if S.SendVerificationEmail() {
				G.Msg.Success = "Email Has Been Sent Successfully."
			}
		} else {
			G.Msg.Success = "Already Activated."
		}
	}
	c.Redirect(http.StatusFound, "/login")
}


func ActivateAccount(c * gin.Context) {
	encEmail := c.Param("encEmail")
	emailVf := c.Param("emailVf")
	var err error
	var decoded []byte
	decoded, err = base64.URLEncoding.DecodeString(encEmail)
	if err != nil {
		log.Println("AuthController.go Log1", err.Error())
		c.HTML(http.StatusOK, "404.html",nil)
		return
	}

	G.User.Email = string(decoded)
	G.User.EmailVerification.String = emailVf
	var success bool

	G.User, success = R.ActivateAccount(G.User)
	if success {
		G.Msg.Success = "Congratulations, Your Account Is Activated."
		c.Redirect(http.StatusFound, "/login")
	} else {
		c.HTML(http.StatusOK, "404.html",nil)
	}
}


func LoginGet(c *gin.Context) {
	if M.IsGuest(c, store, sc) {
		c.HTML(http.StatusOK, "login.html", G.Msg)
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
	return
}


func LoginPost(c *gin.Context) {
	G.User.Email = c.PostForm("email")
	password := c.PostForm("password")
	rememberMe, _ := strconv.Atoi(c.PostForm("remember_me"))
	var success bool
	G.User, success = R.Login(G.User)
	if success {
		err := bcrypt.CompareHashAndPassword([]byte(G.User.Password), []byte(password))
		if err != nil {
			G.Msg.Fail = "Wrong Credentials."
			c.Redirect(http.StatusFound, "/login")
		} else {
			if G.User.ActiveStatus == 1 {
				G.User.RememberToken.String = H.RandomString(60)
				G.User.RememberToken.Valid = true

				if !R.SetRememberToken(G.User) {
					G.Msg.Fail = "Some Internal Server Error Occurred. Please Try Again."
					c.Redirect(http.StatusFound, "/login")
					return
				}

				session, _ := store.Get(c.Request, "login_token")
				session.Values["userEmail"] = G.User.Email
				session.Values["remember_token"] = G.User.RememberToken.String
				session.Save(c.Request, c.Writer)

				if rememberMe == 1 {
					session.Options.MaxAge = 60 * 60 * 24 * 365
					session.Save(c.Request, c.Writer)
				}
				if G.User.RoleID == 1 {
					c.Redirect(http.StatusFound, "/dashboard")
				} else if G.User.RoleID == 2 {
					c.Redirect(http.StatusFound, "/home")
				}
			} else if G.User.ActiveStatus == 2 {
				if G.Msg.Fail == "" {
					G.Msg.Fail = "You Are Suspended. Contact With The Authority Quickly."
				}
				c.Redirect(http.StatusFound, "/login")
			} else {
				var link template.HTML
				link = "<a href='http://localhost:2000/resend-email-verification'>Click Here To Send Verification Email</a>"
				if G.Msg.Fail == "" {
					G.Msg.Fail = "Please Activate Your Account, " + link + "."
				}
				c.Redirect(http.StatusFound, "/login")
			}
		}

	} else {
		if G.Msg.Fail == "" {
			G.Msg.Fail = "User Not Found."
		}
		c.Redirect(http.StatusFound, "/login")
	}
}


func ForgotPassword(c *gin.Context) {
	if M.IsGuest(c, store, sc) {
		c.HTML(http.StatusOK, "forgot-password.html", G.Msg)
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
}


func SendPasswordResetLink(c *gin.Context) {
	var success bool
	G.User.Email = c.PostForm("email")
	G.User, success = R.ReadWithEmail(G.User)
	if !success {
		G.Msg.Fail = "User Not Found With This Email."
		c.Redirect(http.StatusFound, "/forgot-password")
		return
	}
	G.PS.Email = G.User.Email
	G.PS.Token.String = H.RandomString(60)
	G.PS.Token.Valid = true
	if !R.SendPasswordResetLink(G.PS) {
		return
	}
	if S.SendPasswordResetLinkEmail() {
		G.Msg.Success = "Reset Password Link Sent Successfully. Check Your Email."
	}
	c.Redirect(http.StatusFound, "/login")
}


func ResetPasswordGet(c *gin.Context) {
	encEmail := c.Param("email")
	var err error
	var decoded []byte
	decoded, err = base64.URLEncoding.DecodeString(encEmail)
	if err != nil {
		log.Println("AuthController.go Log2", err.Error())
		c.HTML(http.StatusOK, "404.html",nil)
		return
	}

	G.PS.Email = string(decoded)
	G.PS.Token.String = c.Param("token")
	G.PS.Token.Valid = true
	if !R.ResetPasswordGet(G.PS) {
		c.HTML(http.StatusOK, "404.html", nil)
		return
	}
	c.HTML(http.StatusOK, "reset-password.html", G.Msg)
}


func ResetPasswordPost(c *gin.Context) {
	password := c.PostForm("password")
	confirmPass := c.PostForm("confirm-password")
	if password != confirmPass {
		G.Msg.Fail = "Confirm Password Doesn't Match."
		encEmail := base64.URLEncoding.EncodeToString([]byte(G.PS.Email))
		c.Redirect(http.StatusFound, "/reset-password/"+encEmail+"/"+G.PS.Token.String)
		return
	}
	cost := bcrypt.DefaultCost
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	G.User.Password = string(hash)
	G.User.Email = G.PS.Email
	if !R.ResetPasswordPost(G.User, G.PS) {
		G.Msg.Fail = "Some Internal Server Error Occurred, Please Try Again Later."
		encEmail := base64.URLEncoding.EncodeToString([]byte(G.PS.Email))
		c.Redirect(http.StatusFound, "/reset-password/"+encEmail+"/"+G.PS.Token.String)
		return
	}
	G.Msg.Success = "Your Password Is Reset Successfully."
	c.Redirect(http.StatusFound, "/login")
}


func Home(c *gin.Context) {
	if M.IsAuthUser(c, store, sc) {
		c.HTML(http.StatusOK, "home.html", G.User)
	}
	return
}


func Dashboard(c *gin.Context) {
	if M.IsAuthAdminUser(c, store, sc) {
		c.HTML(http.StatusOK, "dashboard.html", G.User)
	}
	return
}


func Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "login_token")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	R.Logout(G.User)

	var user Models.User
	G.User = user

	c.Redirect(http.StatusFound, "/login")
}
