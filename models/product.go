package model

type Product struct {
	ID              int64  `json:"id"`
	NamaProduk      string `json:"nama_produk"`
	DeskripsiProduk string `json:"deskripsi_produk"`
	GambarProduk    string `json:"gambar_produk"`
	KategoriID      int64  `json:"kategori_id"`
	StokProduk      int    `json:"stok_produk"`
}
