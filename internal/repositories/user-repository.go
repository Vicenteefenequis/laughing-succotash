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
		err := r.db.Model(&_user).Updates(u).Error

		if err != nil {
			return &domain.User{}, errors.New(apperrors.IllegalOperation, err, "failed", "fail to update user on database")
		}

		return &u, nil
	}

	err := r.db.Create(&u).Error

	if err != nil {
		return &domain.User{}, errors.New(apperrors.IllegalOperation, err, "failed", "failed save user on database")
	}

	return &u, nil

}
func (r *User) Delete(id string) error {
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
