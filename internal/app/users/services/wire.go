// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/Infinity-OJ/Server/internal/app/users/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateUsersService(cf string, sto repositories.UsersRepository) (UsersService, error) {
	panic(wire.Build(testProviderSet))
}
