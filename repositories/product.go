package repositories

import (
	"database/sql"

	"github.com/aldysp34/deeptech-test/config"
	model "github.com/aldysp34/deeptech-test/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository() *ProductRepository {
	db := config.Connect()
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	query := `INSERT INTO products (nama_produk, deskripsi_produk, gambar_produk, kategori_id, stok_produk) VALUES (?, ?, ?, ?, ?)`
	result, err := r.DB.Exec(query, product.NamaProduk, product.DeskripsiProduk, product.GambarProduk, product.KategoriID, product.StokProduk)
	if err != nil {
		return err
	}
	product.ID, _ = result.LastInsertId()
	return nil
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.DB.Query(`SELECT id, nama_produk, deskripsi_produk, gambar_produk, kategori_id, stok_produk FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.NamaProduk, &p.DeskripsiProduk, &p.GambarProduk, &p.KategoriID, &p.StokProduk)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int64) (*model.Product, error) {
	var p model.Product
	err := r.DB.QueryRow(`SELECT id, nama_produk, deskripsi_produk, gambar_produk, kategori_id, stok_produk FROM products WHERE id = ?`, id).
		Scan(&p.ID, &p.NamaProduk, &p.DeskripsiProduk, &p.GambarProduk, &p.KategoriID, &p.StokProduk)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(p *model.Product) error {
	_, err := r.DB.Exec(`UPDATE products SET nama_produk=?, deskripsi_produk=?, gambar_produk=?, kategori_id=?, stok_produk=? WHERE id=?`,
		p.NamaProduk, p.DeskripsiProduk, p.GambarProduk, p.KategoriID, p.StokProduk, p.ID)
	return err
}

func (r *ProductRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}
