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

func (u *User) Find(ids []string, limit int, offset int) ([]domain.User, error) {
	_users, err := u.repository.Get(ids, limit, offset)

	if err != nil {
		return []domain.User{}, nil
	}

	return _users, nil
}

func (u *User) Delete(id string) error {
	err := u.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
func (u *User) Update(user domain.User) (*domain.User, error) {

	_user, err := u.repository.Save(user)

	if err != nil {
		return &domain.User{}, err
	}

	return _user, nil
}
