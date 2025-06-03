package model

type Category struct {
	ID                      int64  `json:"id"`
	NamaKategoriProduk      string `json:"nama_kategori_produk"`
	DeskripsiKategoriProduk string `json:"deskripsi_kategori_produk"`
}
