package main

import (
	"api_buku/controller"
	"api_buku/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//melakukan koneksi ke db
	dsn := "root:@tcp(127.0.0.1:3306)/api_buku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database")//program akan berhenti
	}
	
	// membuat tabel otomatis dengan migration
	db.AutoMigrate(&model.Buku{})

	bukuRepository := model.NewRepository(db)
	bukuService 	:= model.NewService(bukuRepository)
	bukuController := controller.NewBukuController(bukuService)

	r := gin.Default()
	r.GET("/buku", bukuController.GetBukuController)    //get all data
	r.GET("/buku/:id", bukuController.GetBukuByIdController) //get data by id
	r.POST("/buku", bukuController.PostBukuController)
	r.PUT("/buku/:id", bukuController.UpdateBukuController)
	r.DELETE("/buku/:id", bukuController.DeleteBukuController)

	r.Run()
}