package Migrtaions

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	Mod "BoilerPlateWithAuthInGo/Models"
	"github.com/jinzhu/gorm"
)


func Migrate() {
	db := Cfg.DBConnect()
	db.AutoMigrate(&Mod.Role{})
	db.AutoMigrate(&Mod.User{})
	db.AutoMigrate(&Mod.PasswordReset{})
	AddForeignKeys(db)
	defer db.Close()
}

func AddForeignKeys(db *gorm.DB) {
	db.Model(&Mod.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.PasswordReset{}).AddForeignKey("email", "users(email)", "RESTRICT", "RESTRICT")
}


