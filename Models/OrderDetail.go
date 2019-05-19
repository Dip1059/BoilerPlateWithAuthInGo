package Models

import (
	"github.com/jinzhu/gorm"
)

type OrderDetail struct{
	gorm.Model
	OrderID uint	`gorm:"index; not null"`
	ProductID uint	`gorm:"index; not null"`
	Quantity int	`gorm:"not null"`
	Product Product
}