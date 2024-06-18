package model

type BukuResponse struct {
	Id        int    `json:"id"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
	Rating    int    `json:"rating"`
}