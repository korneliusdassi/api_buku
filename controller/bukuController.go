// mengakses ke service
package controller

import (
	"api_buku/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bukuController struct {
	bukuService model.Service
}

func NewBukuController(bukuService model.Service) *bukuController {
	return &bukuController{bukuService}
}

func responseBuku(buku model.Buku) (model.BukuResponse) {
	return model.BukuResponse{
		Id 			: buku.Id,
		Judul 		: buku.Judul,
		Deskripsi 	: buku.Deskripsi,
		Harga 		: buku.Harga,
		Rating 		: buku.Rating,
	}
}

//method untuk menampilkan semua data buku
func (controller *bukuController) GetBukuController(c *gin.Context) {
	buku, err := controller.bukuService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}

	var buku_response []model.BukuResponse

	for _, buku := range buku {
		bukuResponse := responseBuku(buku)
		buku_response = append(buku_response, bukuResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"status" : "success",
		"data" : buku_response,
	})
}

// method untuk menampilkan data buku by id
func (controller *bukuController) GetBukuByIdController(c *gin.Context) {
	id_string := c.Param("id")

	id, _:= strconv.Atoi(id_string)
	buku, err := controller.bukuService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}
	bukuResponse := responseBuku(buku)

	//cek apakah data yg dicari ada atau tidak
	if bukuResponse.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data tidak ditemukan",
		})
		return
	}
	//jika datanya ada maka tampilkan
	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"status" : "success",
		"data" : bukuResponse,
	})
}

//method untuk menambahkan data buku
func (controller *bukuController) PostBukuController(c *gin.Context) {
	var bukuRequest model.BukuRequest

	err := c.ShouldBindJSON(&bukuRequest)
	if err != nil {
		errorMessages := []string{} //untuk menampilkan banyak pesan error
		for _, error := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("field %s tidak boleh kosong, condition: %s", error.Field(), error.ActualTag())
			errorMessages = append(errorMessages, errMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":errorMessages,
		})
		return
	}
	//panggil service
	newBuku, err := controller.bukuService.Create(bukuRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "Successfully",
		"data":   responseBuku(newBuku),
	})
}

//methode update data
func (controller *bukuController) UpdateBukuController (c *gin.Context){
	var bukuRequest model.BukuRequest

	err := c.ShouldBindJSON(&bukuRequest)
	if err != nil {
		errorMessages := []string{} //untuk menampilkan banyak pesan error
		for _, error := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("field %s tidak boleh kosong, condition: %s", error.Field(), error.ActualTag())
			errorMessages = append(errorMessages, errMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":errorMessages,
		})
		return
	}

	id_string := c.Param("id")
	id, _:= strconv.Atoi(id_string)
	//panggil servis
	newBuku, err := controller.bukuService.Update(id, bukuRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"status": "Successfully",
		"data": responseBuku(newBuku),
	})
}

//method hapus data by id
func (controller *bukuController) DeleteBukuController (c *gin.Context){
	id_string := c.Param("id")
	id, _ := strconv.Atoi(id_string)

	_, err := controller.bukuService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"status": "Successfully delete data",
	})
}

