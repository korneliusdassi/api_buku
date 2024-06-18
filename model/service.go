/*
	servis ini berhubungan dgn logic/fitur
	logic ini berhubungan/mengakses ke repository
*/

package model

type Service interface {
	Create(bukuRequest BukuRequest) (Buku, error)
	FindAll() ([]Buku, error)
	FindById(id int) (Buku, error)
	Update(id int, bukuRequest BukuRequest) (Buku, error)
	Delete(id int) (Buku, error)
}

type service struct {
	repositori Repository
}

func NewService(repositori Repository) *service {
	return &service{repositori}
}

func (s *service) FindAll() ([]Buku, error) {
	buku, err := s.repositori.FindAll()
	return buku, err
}

func (s *service) FindById(id int) (Buku, error) {
	buku, err := s.repositori.FindById(id)
	return buku, err
}

func (s *service) Create(bukuRequest BukuRequest) (Buku, error) {
	buku := Buku{ //kita mapping satu-satu dari bukuRequest ke buku
		Judul 		: bukuRequest.Judul,
		Deskripsi 	: bukuRequest.Deskripsi,
		Harga 		: bukuRequest.Harga,
		Rating 		: bukuRequest.Rating,
	}
	newBuku, err := s.repositori.Create(buku)
	return newBuku, err	
}

func (s *service) Update(id int, bukuRequest BukuRequest) (Buku, error) {
	buku, _ := s.repositori.FindById(id)

	buku.Judul 		= bukuRequest.Judul
	buku.Deskripsi = bukuRequest.Deskripsi
	buku.Harga 		= bukuRequest.Harga
	buku.Rating 	= bukuRequest.Rating

	newBuku, err := s.repositori.Update(buku)
	return newBuku, err
}

func (s *service) Delete(id int) (Buku, error){
	buku, _ := s.repositori.FindById(id)
	newBuku, err := s.repositori.Delete(buku)
	return newBuku, err
}


