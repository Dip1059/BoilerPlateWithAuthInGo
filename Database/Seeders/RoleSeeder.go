package Seeders

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	Mod "BoilerPlateWithAuthInGo/Models"
)

var roles =make([]Mod.Role,0)

func RoleSeeder() {
	role1()
	role2()
	db := Cfg.DBConnect()
	for i,_ := range roles {
		db.FirstOrCreate(&roles[i],&Mod.Role{Name:roles[i].Name})
	}
}

func role1() {
	var role = Mod.Role{
		Name: "Admin",
		Status: 1,
	}
	roles = append(roles, role)
}

func role2() {
	var role = Mod.Role{
		Name: "User",
		Status: 1,
	}
	roles = append(roles, role)
}
