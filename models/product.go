package models

import "time"

type Product struct {
	ID          int       `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Price       float64   `gorm:"size:25;not null" json:"price"`
	Description string    `gorm:"size:255;not null" json:"description"`
	SalePrice   float64   `gorm:"size:25;not null" json:"sale_price"`
	CategoryID  int       `gorm:"size:25;not null" json:"category_id"`
	BrandID     int       `gorm:"size:25;not null" json:"brand_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
