package models

import "gorm.io/gorm"

type Brand struct {
	*gorm.Model
	Name string `json:"name" gorm:"type:varchar(50)"`
	Year int    `json:"year" gorm:"type:int(11)"`
}

type BrandInput struct {
	Name string `json:"name" binding:"required,alpha"`
	Year int    `json:"year" binding:"required"`
}

type Laptop struct {
	*gorm.Model
	Name    string `json:"name" gorm:"type:varchar(50)"`
	Price   int    `json:"price" gorm:"type:int(11)"`
	BrandID int    `json:"brand_id" gorm:"type:int(11)"`
	Brand   Brand  `json:"brand" gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
