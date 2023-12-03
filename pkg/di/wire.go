//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/easy-health/pkg/api"
	handler "github.com/easy-health/pkg/api/handler"
	config "github.com/easy-health/pkg/config"
	db "github.com/easy-health/pkg/db"
	repository "github.com/easy-health/pkg/repository"
	usecase "github.com/easy-health/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewAdminRepository,
		repository.NewUserRepository,
		repository.NewDoctorRepository,
		repository.NewAuthRepository,
		repository.NewCategoriesRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewDoctorUseCase,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewDoctorHandler,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
