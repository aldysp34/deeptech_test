package service

import (
	model "github.com/aldysp34/deeptech-test/models"
	"github.com/aldysp34/deeptech-test/repositories"
)

type CategoryService struct {
	Repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) Create(category *model.Category) error {
	return s.Repo.Create(category)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.Repo.GetAll()
}

func (s *CategoryService) GetByID(id int64) (*model.Category, error) {

	return s.Repo.GetByID(id)
}

func (s *CategoryService) Update(category *model.Category) error {
	return s.Repo.Update(category)
}

func (s *CategoryService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
