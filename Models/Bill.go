package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Bill struct{
	gorm.Model
	UserID uint	`gorm:"index; not null"`
	FullName string	`gorm:"not null"`
	Email sql.NullString
	Address sql.NullString
	Color sql.NullString
	Phone string	`gorm:"not null"`
	Status int	`gorm:"type:tinyint(4); not null"`
}