package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type string

const (
	client = Type("client")
	store  = Type("store")
)

type User struct {
	ID   string   `gorm:"primaryKey;type:char(36)" json:"id"`
	Name string `json:"name"`
	Type Type   `json:"type"`
}

func NewUser(id string, name string, _type Type) User {
	return User{
		ID:   id,
		Name: name,
		Type: _type,
	}
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.ID = generateUid()
	return nil
}

func generateUid() string {
	return uuid.New().String()
}
