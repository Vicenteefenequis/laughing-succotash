package service

import (
	"laughing-succostash/internal/core/domain"
	repository_port "laughing-succostash/internal/core/ports/repositories"
)

type Product struct {
	repository repository_port.Product
}

func NewProductService(repository repository_port.Product) *Product {
	return &Product{
		repository,
	}
}

func (u *Product) Create(product domain.Product) (*domain.Product, error) {

	productSaved, err := u.repository.Save(product)

	if err != nil {
		return nil, err
	}

	return productSaved, nil
}

func (u *Product) Find(ids []string, limit int, offset int) ([]domain.Product, error) {
	_products, err := u.repository.Get(ids, limit, offset)

	if err != nil {
		return []domain.Product{}, nil
	}

	return _products, nil
}

func (u *Product) Delete(id string) error {
	err := u.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
func (u *Product) Update(product domain.Product) (*domain.Product, error) {

	_product, err := u.repository.Save(product)

	if err != nil {
		return &domain.Product{}, err
	}

	return _product, nil
}
