package service_port

import "laughing-succostash/internal/core/domain"

type Category interface {
	Create(category domain.Category) (*domain.Category, error)
	Find(ids []string, limit int, offset int) ([]domain.Category, error)
	Delete(id string) error
	Update(Category domain.Category) (*domain.Category, error)
}
