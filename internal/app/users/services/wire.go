// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server/internal/app/users/repositories"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
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
