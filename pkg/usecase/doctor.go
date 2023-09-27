package usecase

import (
	"context"
	"errors"

	interfaces "github.com/easy-health/pkg/repository/interface"
	services "github.com/easy-health/pkg/usecase/interface"
	"github.com/easy-health/pkg/utils/req"
	"golang.org/x/crypto/bcrypt"
)

type DoctorUseCase struct {
	doctorRepo interfaces.DoctorRepository
}

func NewDoctorUseCase(doctorRepo interfaces.DoctorRepository) services.DoctorUseCase {
	return &DoctorUseCase{
		doctorRepo: doctorRepo,
	}
}

func (connect *DoctorUseCase) Registration(ctx context.Context, data req.DoctorRegistration) error {
	check, err := connect.doctorRepo.EmailChecking(data.Email)
	if err != nil {
		return err
	}

	if check {
		return errors.New("email is already registered")
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err = connect.doctorRepo.Register(ctx, data, string(hashedPassword)); err != nil {
		return err
	}
	return nil

}
