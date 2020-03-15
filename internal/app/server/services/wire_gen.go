// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateProductsService(cf string, usersClients proto.UsersClient) (ProductsService, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	productsService := NewProductService(logger, usersClients)
	return productsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, ProviderSet)