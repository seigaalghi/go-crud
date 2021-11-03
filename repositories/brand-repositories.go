package repositories

import (
	"github.com/seigaalghi/go-crud/models"
	"gorm.io/gorm"
)

func CreateBrand(db *gorm.DB, input *models.BrandInput) (*models.Brand, error) {
	brand := models.Brand{
		Name: input.Name,
		Year: input.Year,
	}
	err := db.Model(&brand).Create(&brand).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func GetBrands(db *gorm.DB) ([]*models.Brand, error) {
	var brands []*models.Brand
	err := db.Where("year = ?", 1998).Limit(10).Find(&brands).Error
	if err != nil {
		return nil, err
	}
	return brands, nil
}
