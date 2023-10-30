// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/easy-health/pkg/api"
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/config"
	"github.com/easy-health/pkg/db"
	"github.com/easy-health/pkg/repository"
	"github.com/easy-health/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase, cfg)
	doctorRepository := repository.NewDoctorRepository(gormDB)
	adminRepository := repository.NewAdminRepository(gormDB)
	doctorUseCase := usecase.NewDoctorUseCase(doctorRepository, adminRepository)
	doctorHandler := handler.NewDoctorHandler(doctorUseCase)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, doctorRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, doctorHandler, adminHandler)
	return serverHTTP, nil
}
