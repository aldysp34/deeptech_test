package service

import (
	"errors"

	model "github.com/aldysp34/deeptech-test/models"
	"github.com/aldysp34/deeptech-test/repositories"
	"github.com/aldysp34/deeptech-test/utils"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	Repo *repositories.AdminRepository
}

func NewAdminService(repo *repositories.AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) CreateAdmin(req model.CreateAdminRequest) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	admin := model.Admin{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hashed),
		BirthDate: req.BirthDate,
		Gender:    req.Gender,
	}

	return s.Repo.Create(admin)
}

func (s *AdminService) Login(email, password string) (string, error) {
	admin, err := s.Repo.FindByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) != nil {
		return "", errors.New("invalid login")
	}

	return utils.GenerateJWT(admin.ID)
}

func (s *AdminService) UpdateProfile(adminID int64, req model.UpdateProfileRequest) error {
	return s.Repo.Update(adminID, req)
}

func (s *AdminService) ListAdmins() ([]model.Admin, error) {
	return s.Repo.List()
}

func (s *AdminService) GetAdminByID(id int64) (*model.Admin, error) {
	return s.Repo.FindByID(id)
}

func (s *AdminService) DeleteAdmin(id int64) error {
	return s.Repo.Delete(id)
}
