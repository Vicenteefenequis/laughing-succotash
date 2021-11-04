package service_port

import "laughing-succostash/internal/core/domain"

type Category interface {
	Create(category domain.Category) (*domain.Category, error)
	FindAll(ids []string, limit int, offset int) ([]domain.Category, error)
	FindOne(ids []string) (*domain.Category, error)
	Delete(id string) error
	Update(Category domain.Category) (*domain.Category, error)
}
