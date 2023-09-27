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
	query := `SELECT COUNT(*) FROM Doctor_Profiles WHERE email = ?`

	// Execute the query and scan the result into count
	if err := c.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false, err
	}

	// If count is greater than 0, the email exists; otherwise, it doesn't
	return count > 0, nil
}

func (c *DoctorDatabase) Register(ctx context.Context, doctor req.DoctorRegistration, hashedPassword string) error {
	
	doctorProfile := domain.DoctorProfiles{
		Name:          doctor.Name,
		Email:         doctor.Email,
		Password:      string(hashedPassword), // Store the hashed password in the database
		Specialise:    doctor.Specialise,
		LicenseNumber: doctor.LicenseNumber,
		Approved:      false,
	}

	// Insert the doctor's information into the database
	if err := c.DB.Create(&doctorProfile).Error; err != nil {
		return err
	}

	return nil
}
