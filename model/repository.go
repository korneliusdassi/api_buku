/*
repositori ini bertanggung jawab terhadap model untuk menjadi peranta ke database
dari main->controller->service->repository->db->mysql
*/
package model

import "gorm.io/gorm"

//definisikan interface yg merupakan kumpulan dari method utk melakukan CRUD
type Repository interface {
	Create(buku Buku) (Buku, error)
	FindAll() ([]Buku, error) //parameternya kosong, dan mengembalikan buku dan eror
	FindById(id int) (Buku, error) //mengembalikan single buku
	Update(buku Buku) (Buku, error)
	Delete(buku Buku) (Buku, error)
}

//buat struct utk implemntasi interface repository
type RepositoryImpl struct {
	//untuk mengakses ke database
	db *gorm.DB
}

//function untuk instans dari struct RepositoryImpl
func NewRepository(db *gorm.DB) *RepositoryImpl{
	return &RepositoryImpl{db}
}

// function findAll
func(r *RepositoryImpl) FindAll() ([]Buku, error){
	var buku []Buku //merupakan slice dari buku
	err := r.db.Find(&buku).Error
	return buku, err
}

//function findById
func(r *RepositoryImpl) FindById(id int) (Buku, error){
	var buku Buku
	err := r.db.Find(&buku, id).Error
	return buku,err
}

//function Create
func(r *RepositoryImpl) Create(buku Buku) (Buku, error){
	err := r.db.Create(&buku).Error
	return buku,err
}

//function Update
func(r *RepositoryImpl) Update(buku Buku) (Buku, error){
	err := r.db.Save(&buku).Error
	return buku, err
}

//function Delete
func(r *RepositoryImpl) Delete(buku Buku) (Buku, error){
	err := r.db.Delete(&buku).Error
	return buku, err
}