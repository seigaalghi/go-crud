package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/go-crud/models"
	"github.com/seigaalghi/go-crud/repositories"
	"gorm.io/gorm"
)

func CreateBrand(ctx *gin.Context) {
	var input models.BrandInput
	db := ctx.MustGet("db").(*gorm.DB)
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	data, err := repositories.CreateBrand(db, &input)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "succesfully created",
		"data":    data,
	})
}

func GetBrands(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	params := ctx.Query("id")
	fmt.Println("Ini adalah id", params)
	data, err := repositories.GetBrands(db)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "succesfully retrieved the data",
		"data":    data,
	})
}
