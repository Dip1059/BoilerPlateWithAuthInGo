package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type PasswordReset struct {
	gorm.Model
	Email string	`gorm:"index; not null"`
	Token sql.NullString
	Status int	`gorm:"type:tinyint(4); not null"`
}
