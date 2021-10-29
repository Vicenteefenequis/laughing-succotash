package repositories

import (
	"gorm.io/gorm"
	"laughing-succostash/internal/core/domain"
)

type User struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (r *User) Save(u domain.User) (*domain.User, error) {

	var _user domain.User

	r.db.First(&_user,"id = ?", u.ID)

	if _user.ID != "" {
		err := r.db.Model(&_user).Updates(u).Error

		if err != nil {
			return nil, err
		}

		return &u, nil
	}

	err := r.db.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return &u, nil

}
func (r *User) Delete(id string) error {
	return nil
}

func (r *User) Get(id string) (*domain.User, error) {
	return nil, nil
}
