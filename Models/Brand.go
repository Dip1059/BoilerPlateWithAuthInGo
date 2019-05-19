package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Brand struct{
	gorm.Model
	Name string	`gorm:"not null"`
	Description sql.NullString
	Status int	`gorm:"type:tinyint(4); not null"`
}