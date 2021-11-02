package repositories

import (
	"github.com/matiasvarela/errors"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/pkg/apperrors"
)

type Category struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *Category {
	return &Category{
		db: db,
	}
}

func (r *Category) Save(c domain.Category) (*domain.Category, error) {

	var category domain.Category

	r.db.First(&category, "id = ?", c.ID)

	if category.ID != "" {

		tx := r.db.Model(&category).UpdateColumns(&c)

		if tx.Error != nil {
			return &domain.Category{}, errors.New(apperrors.IllegalOperation, tx.Error, "failed", "fail to update user on database")
		}

		return &category, nil
	}

	err := r.db.Create(&c).Error

	if err != nil {
		return &domain.Category{}, errors.New(apperrors.IllegalOperation, err, "failed", "failed save user on database")
	}

	return &c, nil
}

func (r *Category) Delete(id string) error {

	tx := r.db.Where("id", id).Delete(&domain.Category{})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *Category) Get(id string) (*domain.Category, error) {
	var category domain.Category

	tx := r.db.First(&category, "id = ?", id)

	if category.ID != "" {
		return &category, nil
	}

	return &domain.Category{}, errors.New(apperrors.IllegalOperation, tx.Error, "User does not exists", "failed to get user on database")
}

func (r *Category) FindAll() ([]domain.Category, error) {
	var categories []domain.Category

	result := r.db.Find(&categories)

	if result.Error != nil {
		return []domain.Category{}, errors.New(apperrors.EmptyResult, result.Error, "User empty result", "User empty result")
	}

	return categories, nil
}
