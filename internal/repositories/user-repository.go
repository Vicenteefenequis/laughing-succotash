package repositories

import (
	"github.com/matiasvarela/errors"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/pkg/apperrors"
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

	r.db.First(&_user, "id = ?", u.ID)

	if _user.ID != "" {

		tx := r.db.Model(&_user).UpdateColumns(&u)

		if tx.Error != nil {
			return &domain.User{}, errors.New(apperrors.IllegalOperation, tx.Error, "failed", "fail to update user on database")
		}

		return &_user, nil
	}

	err := r.db.Create(&u).Error

	if err != nil {
		return &domain.User{}, errors.New(apperrors.IllegalOperation, err, "failed", "failed save user on database")
	}

	return &u, nil

}
func (r *User) Delete(id string) error {

	tx := r.db.Where("id", id).Delete(&domain.User{})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *User) Get(id string) (*domain.User, error) {
	var user domain.User

	tx := r.db.First(&user, "id = ?", id)

	if user.ID != "" {
		return &user, nil
	}

	return &domain.User{}, errors.New(apperrors.IllegalOperation, tx.Error, "User does not exists", "failed to get user on database")
}

func (r *User) FindAll() ([]domain.User, error) {
	var users []domain.User

	result := r.db.Find(&users)

	if result.Error != nil {
		return []domain.User{}, errors.New(apperrors.EmptyResult, result.Error, "User empty result", "User empty result")
	}

	return users, nil
}
