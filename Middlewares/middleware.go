package Middlewares

import (
	G "BoilerPlateWithAuthInGo/Globals"
	R "BoilerPlateWithAuthInGo/Repositories"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)


func IsGuest(c *gin.Context, store *sessions.CookieStore, sc *securecookie.SecureCookie) bool {
	session, _ := store.Get(c.Request, "login_token")
	email := session.Values["userEmail"]
	rememberToken := session.Values["remember_token"]
	var success bool

	if email != nil && rememberToken != nil {
		G.User.Email = session.Values["userEmail"].(string)
		G.User, success = R.ReadWithEmail(G.User)
		if !success {
			c.Redirect(http.StatusFound, "/logout")
			return false
		}
		if rememberToken.(string) != G.User.RememberToken.String {
			c.Redirect(http.StatusFound, "/logout")
			return false
		}
		if G.User.ActiveStatus == 2 {
			c.Redirect(http.StatusFound, "/logout")
		} else if G.User.ActiveStatus == 1 && G.User.RoleID == 1 {
			c.Redirect(http.StatusFound, "/dashboard")
		} else if G.User.ActiveStatus == 1 && G.User.RoleID == 2 {
			c.Redirect(http.StatusFound, "/home")
		}
		return false
	}
	return true
}


func IsAuthUser(c *gin.Context, store *sessions.CookieStore, sc *securecookie.SecureCookie) bool {
	session, _ := store.Get(c.Request, "login_token")
	email := session.Values["userEmail"]
	rememberToken := session.Values["remember_token"]
	var success bool

	if email != nil && rememberToken != nil {
		G.User.Email = session.Values["userEmail"].(string)
		G.User, success = R.ReadWithEmail(G.User)
		if !success {
			G.Msg.Fail = "User Doesn't Exist Anymore."
			c.Redirect(http.StatusFound, "/logout")
			return false
		}
		if rememberToken.(string) != G.User.RememberToken.String {
			//G.Msg.Fail = "Someone Stole Your Cookie From Your Browser. Please Be Cautious."
			c.Redirect(http.StatusFound, "/logout")
			return false
		}
		if G.User.ActiveStatus == 2 {
			G.Msg.Fail = "You Are Suspended. Contact With The Authority Quickly."
			c.Redirect(http.StatusFound, "/logout")
		} else if G.User.ActiveStatus == 1 && G.User.RoleID == 1 {
			c.Redirect(http.StatusFound, "/dashboard")
		} else if G.User.ActiveStatus == 1 && G.User.RoleID == 2 {
			return true
		}
		return false
	}
	c.Redirect(http.StatusFound, "/login")
	return false
}

func IsAuthAdminUser(c *gin.Context, store *sessions.CookieStore, sc *securecookie.SecureCookie) bool {
	session, _ := store.Get(c.Request, "login_token")
	email := session.Values["userEmail"]
	rememberToken := session.Values["remember_token"]
	var success bool

	if email != nil && rememberToken != nil {
		G.User.Email = session.Values["userEmail"].(string)
		G.User, success = R.ReadWithEmail(G.User)
		if !success {
			G.Msg.Fail = "User Doesn't Exist Anymore."
			c.Redirect(http.StatusFound, "/logout")
			return false
		}
		if rememberToken.(string) != G.User.RememberToken.String {
			//G.Msg.Fail = "Someone Stole Your Cookie From Your Browser. Please Be Cautious."
			c.Redirect(http.StatusFound, "/logout")
			return false
		}
		if G.User.ActiveStatus == 2 {
			G.Msg.Fail = "You Are Suspended. Contact With The Authority Quickly."
			c.Redirect(http.StatusFound, "/logout")
		} else if G.User.ActiveStatus == 1 && G.User.RoleID == 2 {
			c.Redirect(http.StatusFound, "/home")
		} else if G.User.ActiveStatus == 1 && G.User.RoleID == 1 {
			return true
		}
		return false
	}
	c.Redirect(http.StatusFound, "/login")
	return false
}
