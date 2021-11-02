package service

import (
	"laughing-succostash/internal/core/domain"
	repository_port "laughing-succostash/internal/core/ports/repositories"
)

type Category struct {
	repository repository_port.Category
}

func NewCategoryService(repository repository_port.Category) *Category {
	return &Category{
		repository,
	}
}

func (u *Category) Create(category domain.Category) (*domain.Category, error) {

	categorySaved, err := u.repository.Save(category)

	if err != nil {
		return nil, err
	}

	return categorySaved, nil
}

func (u *Category) FindAll() ([]domain.Category, error) {
	_category, err := u.repository.FindAll()

	if err != nil {
		return []domain.Category{}, nil
	}

	return _category, nil
}

func (u *Category) FindOne(id string) (*domain.Category, error) {
	_category, err := u.repository.Get(id)

	if err != nil {
		return &domain.Category{}, err
	}
	return _category, nil
}

func (u *Category) Delete(id string) error {
	err := u.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
func (u *Category) Update(category domain.Category) (*domain.Category, error) {

	_category, err := u.repository.Save(category)

	if err != nil {
		return &domain.Category{}, err
	}

	return _category, nil
}
