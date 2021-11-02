package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	ID   string `gorm:"primaryKey;type:char(36)" json:"id"`
	Name string `json:"name"`
}

func (c *Category) BeforeSave(tx *gorm.DB) error {
	c.ID = generateUid()
	return nil
}
