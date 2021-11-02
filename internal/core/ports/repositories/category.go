package repository_port

import "laughing-succostash/internal/core/domain"

type Category interface {
	Save(category domain.Category) (*domain.Category, error)
	Get(id string) (*domain.Category, error)
	FindAll() ([]domain.Category, error)
	Delete(id string) error
}
