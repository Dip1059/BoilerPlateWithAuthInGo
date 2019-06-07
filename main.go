package main

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	Mig "BoilerPlateWithAuthInGo/Database/Migrations"
	Seed "BoilerPlateWithAuthInGo/Database/Seeders"
	"BoilerPlateWithAuthInGo/Routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	Cfg.Config()
	Mig.Migrate()
	Seed.Seed()
	Routes.Routes()
}
