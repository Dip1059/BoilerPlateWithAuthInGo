package Config

import(
	G "BoilerPlateWithAuthInGo/Globals"
	"github.com/joho/godotenv"
	"os"
)

func EmailConfig() {
	godotenv.Load()
	G.EmailEnv.Host = os.Getenv("MAIL_HOST")
	G.EmailEnv.Port = os.Getenv("MAIL_PORT")
	G.EmailEnv.Username = os.Getenv("MAIL_USERNAME")
	G.EmailEnv.Password = os.Getenv("MAIL_PASSWORD")
}
