package Models

import (
	"github.com/jinzhu/gorm"
)

type Order struct{
	gorm.Model
	UserID uint	`gorm:"index; not null"`
	BillID uint	`gorm:"index; not null"`
	PayMethodID uint	`gorm:"index; not null"`
	Total float64	`gorm:"not null"`
	Status int	`gorm:"type:tinyint(4); not null"`
	User User
	Bill Bill
	PayMethod PayMethod
	OrderDetails []OrderDetail
}