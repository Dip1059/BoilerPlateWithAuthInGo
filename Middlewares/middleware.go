package Middlewares

import (
	G "BoilerPlateWithAuthInGo/Globals"
	Mod "BoilerPlateWithAuthInGo/Models"
	R "BoilerPlateWithAuthInGo/Repositories"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)


func IsGuest(c *gin.Context, store *sessions.FilesystemStore) (Mod.User, bool) {
	var user Mod.User
	session, _ := store.Get(c.Request, "login_token")
	email := session.Values["userEmail"]
	
	var success bool
	if email != nil {
		user.Email = session.Values["userEmail"].(string)
		user, success = R.ReadWithEmail(user)
		if !success {
			c.Redirect(http.StatusFound, "/logout")
			return user, false
		}
		if user.ActiveStatus == 2 {
			c.Redirect(http.StatusFound, "/logout")
		} else if user.ActiveStatus == 1 && user.RoleID == 1 {
			c.Redirect(http.StatusFound, "/dashboard")
		} else if user.ActiveStatus == 1 && user.RoleID == 2 {
			c.Redirect(http.StatusFound, "/home")
		}
		return user, false
	}
	return user, true
}


func IsAuthUser(c *gin.Context, store *sessions.FilesystemStore) (Mod.User, bool) {
	var user Mod.User
	session, _ := store.Get(c.Request, "login_token")
	email := session.Values["userEmail"]
	
	var success bool
	if email != nil {
		user.Email = session.Values["userEmail"].(string)
		user, success = R.ReadWithEmail(user)
		if !success {
			G.Msg.Fail = "User Doesn't Exist Anymore."
			c.Redirect(http.StatusFound, "/logout")
			return user, false
		}
		if user.ActiveStatus == 2 {
			G.Msg.Fail = "You Are Suspended. Contact With The Authority Quickly."
			c.Redirect(http.StatusFound, "/logout")
		} else if user.ActiveStatus == 1 && user.RoleID == 1 {
			c.Redirect(http.StatusFound, "/dashboard")
		} else if user.ActiveStatus == 1 && user.RoleID == 2 {
			return user, true
		}
		return user, false
	}
	c.Redirect(http.StatusFound, "/login")
	return user, false
}


func IsAuthAdminUser(c *gin.Context, store *sessions.FilesystemStore) (Mod.User, bool) {
	var user Mod.User
	session, _ := store.Get(c.Request, "login_token")
	email := session.Values["userEmail"]
	
	var success bool

	if email != nil {
		user.Email = session.Values["userEmail"].(string)
		user, success = R.ReadWithEmail(user)
		if !success {
			G.Msg.Fail = "User Doesn't Exist Anymore."
			c.Redirect(http.StatusFound, "/logout")
			return user, false
		}
		if user.ActiveStatus == 2 {
			G.Msg.Fail = "You Are Suspended. Contact With The Authority Quickly."
			c.Redirect(http.StatusFound, "/logout")
		} else if user.ActiveStatus == 1 && user.RoleID == 2 {
			c.Redirect(http.StatusFound, "/home")
		} else if user.ActiveStatus == 1 && user.RoleID == 1 {
			return user, true
		}
		return user, false
	}
	c.Redirect(http.StatusFound, "/login")
	return user, false
}
