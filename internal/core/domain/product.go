package domain

import "gorm.io/gorm"

type Product struct {
	ID         string    `gorm:"primaryKey;type:char(36)" json:"id"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	CategoryId string    `json:"-"`
	Category   *Category `json:"category,omitempty"`
}

func (p *Product) BeforeSave(tx *gorm.DB) error {
	p.ID = generateUid()
	return nil
}
