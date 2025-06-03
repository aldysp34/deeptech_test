package repositories

import (
	"database/sql"

	"github.com/aldysp34/deeptech-test/config"
	model "github.com/aldysp34/deeptech-test/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	db := config.Connect()
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	query := `INSERT INTO categories (nama_kategori_produk, deskripsi_kategori_produk) VALUES (?, ?)`
	result, err := r.DB.Exec(query, category.NamaKategoriProduk, category.DeskripsiKategoriProduk)
	if err != nil {
		return err
	}
	category.ID, _ = result.LastInsertId()
	return nil
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.DB.Query(`SELECT id, nama_kategori_produk, deskripsi_kategori_produk FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.NamaKategoriProduk, &c.DeskripsiKategoriProduk)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id int64) (*model.Category, error) {
	var c model.Category
	err := r.DB.QueryRow(`SELECT id, nama_kategori_produk, deskripsi_kategori_produk FROM categories WHERE id = ?`, id).
		Scan(&c.ID, &c.NamaKategoriProduk, &c.DeskripsiKategoriProduk)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Update(category *model.Category) error {
	_, err := r.DB.Exec(`UPDATE categories SET nama_kategori_produk = ?, deskripsi_kategori_produk = ? WHERE id = ?`,
		category.NamaKategoriProduk, category.DeskripsiKategoriProduk, category.ID)
	return err
}

func (r *CategoryRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
