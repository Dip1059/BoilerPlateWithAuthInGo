package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type PayMethod struct{
	gorm.Model
	Method string	`gorm:"not null"`
	Description sql.NullString
	Status int	`gorm:"type:tinyint(4); not null"`
}