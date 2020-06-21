// +build wireinject

package controllers

import (
	"github.com/infinity-oj/server/internal/app/users/repositories"
	"github.com/infinity-oj/server/internal/app/users/services"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
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
