package Migrtaions

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	G "BoilerPlateWithAuthInGo/Globals"
	Mod "BoilerPlateWithAuthInGo/Models"
)


func Migrate() {
	Cfg.DBConnect()
	G.DB.AutoMigrate(&Mod.Role{})
	G.DB.AutoMigrate(&Mod.User{})
	G.DB.AutoMigrate(&Mod.PasswordReset{})
	AddForeignKeys()
	defer G.DB.Close()
}

func AddForeignKeys() {
	G.DB.Model(&Mod.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	G.DB.Model(&Mod.PasswordReset{}).AddForeignKey("email", "users(email)", "RESTRICT", "RESTRICT")
}


