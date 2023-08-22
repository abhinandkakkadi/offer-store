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
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase,handler.NewUserHandler,http.NewServerHTTP)
	return &http.ServerHTTP{}, nil
}
