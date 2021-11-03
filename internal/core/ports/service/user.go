package service_port

import (
	"laughing-succostash/internal/core/domain"
)

type User interface {
	Create(user domain.User) (*domain.User, error)
	FindAll(ids []string) ([]domain.User, error)
	FindOne(id string) (*domain.User, error)
	Delete(id string) error
	Update(user domain.User) (*domain.User, error)
}
