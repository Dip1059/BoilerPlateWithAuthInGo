package Config

import (
	G "BoilerPlateWithAuthInGo/Globals"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func CreateMessage(key string, value string, c *gin.Context) {
	session, _ := G.Store.Get(c.Request, "message")
	session.Values[key] = value
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetMessage(c *gin.Context) G.Message {
	var msg G.Message
	session, _ := G.Store.Get(c.Request, "message")
	success := session.Values["success"]
	if success != nil {
		msg.Success = template.HTML(success.(string))
	} else {
		msg.Success = ""
	}

	fail := session.Values["fail"]
	if fail != nil {
		msg.Fail = template.HTML(fail.(string))
	} else {
		msg.Fail = ""
	}
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		fmt.Println(err.Error())
	}
	return msg
}

func GetAndSetMessage(key string, value string, c *gin.Context) {
	session, _ := G.Store.Get(c.Request, "message")
	preValue := session.Values[key]
	if preValue == nil {
		session.Values[key] = value
		err := session.Save(c.Request, c.Writer)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}