package Seeders

import(
	Cfg "BoilerPlateWithAuthInGo/Config"
	G "BoilerPlateWithAuthInGo/Globals"
)

func Seed() {
	Cfg.DBConnect()
	RoleSeeder()
	UserSeeder()
	defer G.DB.Close()
}

