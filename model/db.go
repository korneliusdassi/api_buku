package model

import "time"

type Buku struct { // kumpulan dari field di tabel buku
	Id				int
	Judul			string
	Deskripsi	string
	Harga			int
	Rating		int
	CreatedAt 	time.Time
	UpdataedAt  time.Time
}