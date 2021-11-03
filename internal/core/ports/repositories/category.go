package repository_port

import "laughing-succostash/internal/core/domain"

type Category interface {
	Save(category domain.Category) (*domain.Category, error)
	Get(ids []string) ([]domain.Category, error)
	Delete(id string) error
}
