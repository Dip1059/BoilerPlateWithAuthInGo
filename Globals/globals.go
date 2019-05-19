package Globals

import (
	Mod "BoilerPlateWithAuthInGo/Models"
	"github.com/jinzhu/gorm"
	"html/template"
)

type DB_ENV struct {
	Host, Port, Dialect, Username, Password, DBname string
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
	DB *gorm.DB
	Role Mod.Role
	User Mod.User
	PS Mod.PasswordReset
	Msg Message
)


