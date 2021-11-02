package repository_port

import (
	"laughing-succostash/internal/core/domain"
)

type User interface {
	Save(user domain.User) (*domain.User, error)
	Get(id string) (*domain.User, error)
	FindAll() ([]domain.User, error)
	Delete(id string) error
}
