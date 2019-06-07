package Repositories

import (
	Cfg "BoilerPlateWithAuthInGo/Config"
	M "BoilerPlateWithAuthInGo/Models"
)


func ReadWithEmail(user M.User) (M.User, bool) {
	db := Cfg.DBConnect()
	notFound := db.First(&user, "email=?", user.Email).RecordNotFound()

	if notFound {
		defer db.Close()
		return user, false
	} else {
		defer db.Close()
		return user, true
	}
}

func Register(user M.User) (M.User, bool) {
	db := Cfg.DBConnect()
	err := db.Create(&user).Error
	if err != nil {
		defer db.Close()
		return user, false
	}
	defer db.Close()
	return user, true
}

func Login(user M.User) (M.User, bool) {
	var success bool
	user, success = ReadWithEmail(user)
	if success {
		return user, true
	} else {
		return user, false
	}
}


func SetRememberToken(user M.User) bool {
	db := Cfg.DBConnect()

	if db.Model(&user).Where("email=?", user.Email).Updates(
		map[string]interface{}{"remember_token":user.RememberToken.String}).RowsAffected == 0 {
		defer db.Close()
		return false
	}

	defer db.Close()
	return true
}


func ActivateAccount(user M.User) (M.User, bool) {
	db := Cfg.DBConnect()
	user.EmailVerification.Valid = false

	if db.Model(&user).Where("email=? and email_verification=?", user.Email, user.EmailVerification.String).Updates(
		map[string]interface{}{"active_status":1, "email_verification":user.EmailVerification}).RowsAffected == 0 {
		defer db.Close()
		return user, false
	}

	defer db.Close()
	return user, true
}


func SendPasswordResetLink(ps M.PasswordReset) bool {
	db := Cfg.DBConnect()
	err := db.Create(&ps).Error
	if err != nil {
		defer db.Close()
		return false
	}
	defer db.Close()
	return true
}


func ResetPasswordGet(ps M.PasswordReset) bool {
	db := Cfg.DBConnect()

	notFound := db.First(&ps,"email=? and token=? and status=0", ps.Email, ps.Token).RecordNotFound()
	if notFound {
		defer db.Close()
		return false
	}

	defer db.Close()
	return true
}


func ResetPasswordPost(user M.User, ps M.PasswordReset) bool {
	db := Cfg.DBConnect()
	tx := db.Begin()
	err := tx.Model(&user).Where("email=?", user.Email).Updates(
		map[string]interface{}{"password":user.Password}).Error
	if err != nil {
		//log.Println("AuthRepo.go Log8", err.Error())
		tx.Rollback()
		defer db.Close()
		return false
	}

	err = tx.Model(&ps).Where("email=? and token=?", ps.Email, ps.Token.String).Updates(
		map[string]interface{}{"status":1}).Error
	if err != nil {
		//log.Println("AuthRepo.go Log8", err.Error())
		tx.Rollback()
		defer db.Close()
		return false
	}

	ps.Token.Valid = false
	err = tx.Model(&ps).Where("email=?", ps.Email).Updates(
		map[string]interface{}{"token":ps.Token}).Error
	if err != nil {
		//log.Println("AuthRepo.go Log8", err.Error())
		tx.Rollback()
		defer db.Close()
		return false
	}

	err = tx.Commit().Error
	if err != nil {
		//log.Println("AuthRepo.go Log8", err.Error())
		tx.Rollback()
		defer db.Close()
		return false
	}
	defer db.Close()
	return true
}


func Logout(user M.User) {
	db := Cfg.DBConnect()
	user.RememberToken.Valid = false
	db.Model(&user).Where("email=?", user.Email).Update(
		"remember_token", user.RememberToken)
	defer db.Close()

}
