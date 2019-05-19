package Services

import (
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

type Email_Env struct {
	Host, Port, Username, Password string
}

var (
	emailEnv Email_Env
)


func init() {
	godotenv.Load()
	emailEnv = Email_Env{
		Host: os.Getenv("MAIL_HOST"),
		Port: os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}


func SendEmail(from, to, subject, htmlString string) bool {
	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", htmlString)

	port,_ := strconv.Atoi(emailEnv.Port)

	dialer := gomail.NewDialer(emailEnv.Host, port, emailEnv.Username, emailEnv.Password)
	if err := dialer.DialAndSend(mail); err != nil {
		log.Println("EmailService.go Log1", err.Error())
		return false
	}
	return true
}