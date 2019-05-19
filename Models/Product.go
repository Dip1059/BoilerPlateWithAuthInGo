package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Product struct{
	gorm.Model
	CategoryID uint	`gorm:"index; not null"`
	BrandID uint	`gorm:"index; not null"`
	Name string	`gorm:"not null"`
	Price float64	`gorm:"not null"`
	Size sql.NullString
	Color sql.NullString
	Description sql.NullString
	Status int	`gorm:"type:tinyint(4); not null"`
	ImgUrl sql.NullString
	Category Category
	Brand Brand
}