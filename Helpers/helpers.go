package Helpers

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"html/template"
	mrand "math/rand"
	"net/http"
	"time"
)

var Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
var lenLetters = len(Letters)

func RandomString(n int) string {
	mrand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = Letters[mrand.Intn(lenLetters)]
	}
	return string(b)
}


func ParseTemplate(fileName string, templateData interface{}) (string, error) {
	var str string
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return str, err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, templateData); err != nil {
		return str, err
	}
	str = buffer.String()
	return str, nil
}


func NullStringProcess(data sql.NullString) sql.NullString{
	if data.String != "" {
		data.Valid = true
	} else {
		data.Valid = false
	}
	return data
}

func SetCookie(hashKey interface{}, blockKey interface{}, value string, name string, age int, c *gin.Context) {
	var sc *securecookie.SecureCookie
	if blockKey != nil {
		sc = securecookie.New([]byte(hashKey.(string)), []byte(blockKey.(string)))
	} else {
		sc = securecookie.New([]byte(hashKey.(string)), nil)
	}
	encoded, err := sc.Encode(name, value)
	if err != nil {
		fmt.Println(err.Error())
	}

	cookie := http.Cookie{
		Name:     name,
		Value:    encoded,
		MaxAge:   age,
	}
	http.SetCookie(c.Writer, &cookie)
}

func GetCookie(hashKey interface{}, blockKey interface{}, name string, c *gin.Context) string{
	var sc *securecookie.SecureCookie
	if blockKey != nil {
		sc = securecookie.New([]byte(hashKey.(string)), []byte(blockKey.(string)))
	} else {
		sc = securecookie.New([]byte(hashKey.(string)), nil)
	}
	if cookie, err := c.Request.Cookie(name); err == nil {
		var value string
		if err = sc.Decode(name, cookie.Value, &value); err == nil {
			return value
		}
	}
	return ""
}
