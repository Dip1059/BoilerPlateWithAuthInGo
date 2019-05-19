package Helpers

import (
	"bytes"
	"html/template"
	mrand "math/rand"
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