package repositories

import (
	"github.com/matiasvarela/errors"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/pkg/apperrors"
)

type Product struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *Product {
	return &Product{
		db: db,
	}
}

func (r *Product) Save(c domain.Product) (*domain.Product, error) {

	var category domain.Product

	r.db.First(&category, "id = ?", c.ID)

	if category.ID != "" {

		tx := r.db.Model(&category).UpdateColumns(&c)

		if tx.Error != nil {
			return &domain.Product{}, errors.New(apperrors.IllegalOperation, tx.Error, "failed", "fail to update user on database")
		}

		return &category, nil
	}

	err := r.db.Create(&c).Error

	if err != nil {
		return &domain.Product{}, errors.New(apperrors.IllegalOperation, err, "failed", "failed save user on database")
	}

	return &c, nil
}

func (r *Product) Delete(id string) error {

	tx := r.db.Where("id", id).Delete(&domain.Product{})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *Product) Get(ids []string, limit int, offset int) ([]domain.Product, error) {
	var products []domain.Product

	if len(ids) == 0 {
		result := r.db.Debug().Offset(offset).Limit(limit).Preload("Category").Find(&products)
		if result.Error != nil {
			return []domain.Product{}, errors.New(apperrors.EmptyResult, result.Error, "User empty result", "User empty result")
		}
		return products, nil
	}

	tx := r.db.Preload("Category").Find(&products, ids)

	if len(products) != 0 {
		return products, nil
	}

	return []domain.Product{}, errors.New(apperrors.IllegalOperation, tx.Error, "User does not exists", "failed to get user on database")
}

func (r *Product) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	result := r.db.Find(&products)

	if result.Error != nil {
		return []domain.Product{}, errors.New(apperrors.EmptyResult, result.Error, "User empty result", "User empty result")
	}

	return products, nil
}
