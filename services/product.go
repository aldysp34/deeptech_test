package service

import (
	model "github.com/aldysp34/deeptech-test/models"
	"github.com/aldysp34/deeptech-test/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {

	return &ProductService{Repo: repo}
}

func (s *ProductService) Create(p *model.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) List() ([]model.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) GetByID(id int64) (*model.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) Update(p *model.Product) error {
	return s.Repo.Update(p)
}

func (s *ProductService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
