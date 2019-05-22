package Models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"not null; unique_index"`
	Status int	`gorm:"type:tinyint(4); not null"`
	Users []User	`gorm:"save_associations:false; association_save_reference:false"`
}
