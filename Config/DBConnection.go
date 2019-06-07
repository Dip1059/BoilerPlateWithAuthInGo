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

func DBConnect() *gorm.DB{
	var db *gorm.DB
	var err error
	if G.DBEnv.Dialect == "mysql" {
		db, err = gorm.Open(G.DBEnv.Dialect, G.DBEnv.Username+":"+G.DBEnv.Password+"@tcp("+G.DBEnv.Host+":"+
			G.DBEnv.Port+")/"+ G.DBEnv.DBname+"?charset=utf8&parseTime=True&loc=Local")

	} else if G.DBEnv.Dialect == "postgres" {
		db, err = gorm.Open(G.DBEnv.Dialect, "host="+G.DBEnv.Host+" port="+
			G.DBEnv.Port+" user="+G.DBEnv.Username+" dbname="+ G.DBEnv.DBname+" password="+G.DBEnv.Password)
	}
	if err !=nil {
		log.Println("log", err.Error())
	}
	return db
}