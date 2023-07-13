//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandkakkadi/offer-store/pkg/api"
	handler "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	config "github.com/abhinandkakkadi/offer-store/pkg/config"
	db "github.com/abhinandkakkadi/offer-store/pkg/db"
	repository "github.com/abhinandkakkadi/offer-store/pkg/repository"
	usecase "github.com/abhinandkakkadi/offer-store/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP, repository.NewProductRepository, usecase.NewProductUseCase, handler.NewProductHandler, handler.NewOtpHandler, usecase.NewOtpUseCase, repository.NewOtpRepository, repository.NewAdminRepository, usecase.NewAdminUseCase, handler.NewAdminHandler, handler.NewCartHandler, usecase.NewCartUseCase, repository.NewCartRepository)
	return &http.ServerHTTP{}, nil
}
