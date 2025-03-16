package models
import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	OrderID   uint           `json:"order_id" gorm:"not null"`
	Order     Order          `json:"order" gorm:"foreignKey:OrderID"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}