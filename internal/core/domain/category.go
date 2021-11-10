package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	ID       string    `gorm:"primaryKey;type:char(36)" json:"id"`
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryId" json:"products,omitempty"`
}

func (c *Category) BeforeSave(tx *gorm.DB) error {
	c.ID = generateUid()
	return nil
}
