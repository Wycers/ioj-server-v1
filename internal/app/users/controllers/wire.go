// +build wireinject

package controllers

import (
	"github.com/Infinity-OJ/Server/internal/app/users/repositories"
	"github.com/Infinity-OJ/Server/internal/app/users/services"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)

func CreateUsersController(cf string, sto repositories.UsersRepository) (*UsersController, error) {
	panic(wire.Build(testProviderSet))
}
