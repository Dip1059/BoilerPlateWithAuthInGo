package Config

import(
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	G "BoilerPlateWithAuthInGo/Globals"
	"log"
	"os"
)

func init() {
	godotenv.Load()
	G.DBEnv = G.DB_ENV{
		Host:os.Getenv("DB_HOST"),
		Port:os.Getenv("DB_PORT"),
		Dialect:os.Getenv("DB_DIALECT"),
		Username:os.Getenv("DB_USERNAME"),
		Password:os.Getenv("DB_PASSWORD"),
		DBname:os.Getenv("DB_NAME"),
	}
}

func DBConnect() {
	var err error
	G.DB, err = gorm.Open(G.DBEnv.Dialect, G.DBEnv.Username+":"+G.DBEnv.Password+"@tcp("+G.DBEnv.Host+":"+G.DBEnv.Port+")/"+G.DBEnv.DBname+"?parseTime=true")
	if err !=nil {
		log.Println("log", err.Error())
	}
}