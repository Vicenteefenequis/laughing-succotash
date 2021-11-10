package service_port

import "laughing-succostash/internal/core/domain"

type Product interface {
	Create(product domain.Product) (*domain.Product, error)
	Find(ids []string, limit int, offset int) ([]domain.Product, error)
	Delete(id string) error
	Update(product domain.Product) (*domain.Product, error)
}
