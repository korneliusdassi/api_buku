package model

type BukuRequest struct {
	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Harga     int    `json:"harga" binding:"required,number"`
	Rating    int    `json:"rating" binding:"required,number"`
}