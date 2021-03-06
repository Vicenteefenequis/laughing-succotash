package service_port

import (
	"laughing-succostash/internal/core/domain"
)

type User interface {
	Create(user domain.User) (*domain.User, error)
	Find(ids []string, limit int, offset int) ([]domain.User, error)
	Delete(id string) error
	Update(user domain.User) (*domain.User, error)
}
