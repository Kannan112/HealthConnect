package repository

import (
	"context"

	"github.com/easy-health/pkg/domain"
	interfaces "github.com/easy-health/pkg/repository/interface"
	"github.com/easy-health/pkg/utils/req"
	"gorm.io/gorm"
)

type DoctorDatabase struct {
	DB *gorm.DB
}

func NewDoctorRepository(DB *gorm.DB) interfaces.DoctorRepository {
	return &DoctorDatabase{DB}
}

func (c *DoctorDatabase) EmailChecking(email string) (bool, error) {
	var count int64
	query := `SELECT COUNT(*) FROM Doctors WHERE email = ?`

	// Execute the query and scan the result into count
	if err := c.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false, err
	}

	// If count is greater than 0, the email exists; otherwise, it doesn't
	return count > 0, nil
}

func (c *DoctorDatabase) Register(ctx context.Context, doctor req.DoctorRegistration, hashedPassword string, categoryId uint) error {

	doctorProfile := domain.Doctors{
		Name:          doctor.Name,
		Email:         doctor.Email,
		Password:      string(hashedPassword), // Store the hashed password in the database
		About:         doctor.About,           //change it upated is on way
		LicenseNumber: doctor.LicenseNumber,
		CategoriesId:  categoryId,
	}

	// Insert the doctor's information into the database
	if err := c.DB.Create(&doctorProfile).Error; err != nil {
		return err
	}

	return nil
}

func (c *DoctorDatabase) Login(ctx context.Context, data req.DoctorLogin) (domain.Doctors, error) {
	var doctordata domain.Doctors
	if err := c.DB.Raw("select * from doctors where email=$1", data.Email).Scan(&doctordata).Error; err != nil {
		return doctordata, err
	}
	return doctordata, nil
}

func (c *DoctorDatabase) Profile(ctx context.Context, id int) (req.DoctorProfile, error) {
	var doctorProfile req.DoctorProfile
	query := `select * from doctors where id=$1`
	if err := c.DB.Raw(query, id).Scan(&doctorProfile).Error; err != nil {
		return doctorProfile, err
	}
	return doctorProfile, nil
}

func (c *DoctorDatabase) CheckDoctorId(ctx context.Context, id int) (bool, error) {
	var check bool
	if err := c.DB.Raw(`select exists(select * from doctors where id=$1)`, id).Scan(&check).Error; err != nil {
		return false, err
	}
	return check, nil
}
