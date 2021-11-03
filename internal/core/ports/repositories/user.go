package repository_port

import (
	"laughing-succostash/internal/core/domain"
)

type User interface {
	Save(user domain.User) (*domain.User, error)
	Get(ids []string) ([]domain.User, error)
	Delete(id string) error
}
