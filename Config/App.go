package Config

import(
	G "BoilerPlateWithAuthInGo/Globals"
	"github.com/joho/godotenv"
	"os"
)

func AppConfig() {
	godotenv.Load()
	G.AppEnv.Name = os.Getenv("APP_NAME")
	G.AppEnv.Url = os.Getenv("APP_URL")
}
