package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName          string         `gorm:"not null"`
	Email             string         `gorm:"not null; unique_index"`
	Phone             sql.NullString `gorm:"unique_index"`
	PhoneVerification sql.NullString
	Password          string `gorm:"not null"`
	ActiveStatus      int    `gorm:"type:tinyint(4); not null"`
	RoleID            uint   `gorm:"index; not null"`
	EmailVerification sql.NullString
	Role              Role	`gorm:"save_associations:false; association_save_reference:false"`
}

