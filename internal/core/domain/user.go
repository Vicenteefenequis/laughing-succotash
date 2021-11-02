package domain

import (
	"gorm.io/gorm"
)

type Type string

const (
	client = Type("client")
	store  = Type("store")
)

type User struct {
	ID   string `gorm:"primaryKey;type:char(36)" json:"id"`
	Name string `json:"name"`
	Type Type   `json:"type"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.ID = generateUid()
	return nil
}
