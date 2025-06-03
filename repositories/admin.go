package repositories

import (
	"database/sql"

	"github.com/aldysp34/deeptech-test/config"
	model "github.com/aldysp34/deeptech-test/models"
)

type AdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository() *AdminRepository {
	db := config.Connect()
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) Create(admin model.Admin) error {
	_, err := r.DB.Exec(`
		INSERT INTO admins (first_name, last_name, email, password, birth_date, gender)
		VALUES (?, ?, ?, ?, ?, ?)`,
		admin.FirstName, admin.LastName, admin.Email, admin.Password, admin.BirthDate, admin.Gender)
	return err
}

func (r *AdminRepository) FindByEmail(email string) (model.Admin, error) {
	row := r.DB.QueryRow(`SELECT id, first_name, last_name, email, password, birth_date, gender FROM admins WHERE email = ?`, email)

	var admin model.Admin
	err := row.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Password, &admin.BirthDate, &admin.Gender)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *AdminRepository) FindByID(id int64) (*model.Admin, error) {
	row := r.DB.QueryRow(`SELECT id, first_name, last_name, email, password, birth_date, gender FROM admins WHERE id = ?`, id)

	var admin model.Admin
	err := row.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Password, &admin.BirthDate, &admin.Gender)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) Update(id int64, req model.UpdateProfileRequest) error {
	_, err := r.DB.Exec(`
		UPDATE admins SET first_name = ?, last_name = ?, birth_date = ?, gender = ?
		WHERE id = ?`, req.FirstName, req.LastName, req.BirthDate, req.Gender, id)
	return err
}

func (r *AdminRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM admins WHERE id = ?`, id)
	return err
}

func (r *AdminRepository) List() ([]model.Admin, error) {
	rows, err := r.DB.Query(`SELECT id, first_name, last_name, email, birth_date, gender FROM admins`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []model.Admin
	for rows.Next() {
		var a model.Admin
		err := rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Email, &a.BirthDate, &a.Gender)
		if err != nil {
			return nil, err
		}
		admins = append(admins, a)
	}
	return admins, nil
}
