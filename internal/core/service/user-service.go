package service

import (
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/internal/core/ports/repositories"
)

type User struct {
	repository repository_port.User
}

func NewUserService(repository repository_port.User) *User {
	return &User{
		repository,
	}
}

func (u *User) Create(user domain.User) (*domain.User, error) {

	userSaved, err := u.repository.Save(user)

	if err != nil {
		return nil, err
	}

	return userSaved, nil
}

func (u *User) FindAll() ([]domain.User, error) {
	return nil, nil
}

func (u *User) FindOne(id string) (*domain.User, error) {
	return nil, nil
}

func (u *User) Delete(id string) error {
	return nil

}
func (u *User) Update(user domain.User, id string) (*domain.User, error) {
	return nil, nil
}
