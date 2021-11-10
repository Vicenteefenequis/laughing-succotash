package repository_port

import "laughing-succostash/internal/core/domain"

type Product interface {
	Save(product domain.Product) (*domain.Product, error)
	Get(ids []string, limit int, offset int) ([]domain.Product, error)
	Delete(id string) error
}
