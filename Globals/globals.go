package Globals

import (
	Mod "BoilerPlateWithAuthInGo/Models"
	"github.com/gorilla/sessions"
	"html/template"
)

type DB_ENV struct {
	Host, Port, Dialect, Username, Password, DBname string
}

type App_Env struct {
	Name, Url string
}

type Email_Env struct {
	Host, Port, Username, Password string
}

type Message struct {
	Success template.HTML
	Fail template.HTML
}

type EmailGenerals struct {
	From, To, Subject, HtmlString string
}

type UserDataForEmail struct {
	EncEmail string
	User Mod.User
	PS Mod.PasswordReset
}

var(
	DBEnv DB_ENV
	Msg Message
	AppEnv App_Env
	EmailEnv Email_Env
	Store = sessions.NewCookieStore([]byte("secret"))
)


