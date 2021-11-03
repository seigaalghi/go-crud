package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/seigaalghi/go-crud/controllers"
	"github.com/seigaalghi/go-crud/database"
	"github.com/seigaalghi/go-crud/models"
)

func main() {
	server := gin.Default()
	db, err := database.ConnectMysql()
	db.AutoMigrate(&models.Brand{})
	db.AutoMigrate(&models.Laptop{})
	if err != nil {
		fmt.Println(err.Error())
	}

	server.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	serverV1 := server.Group("/api/v1")
	serverV1.POST("/brand", controllers.CreateBrand)
	serverV1.GET("/brand", controllers.GetBrands)
	server.Run(":8000")
}
