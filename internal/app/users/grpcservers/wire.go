// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/Infinity-OJ/Server/internal/app/users/services"
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

func CreateUsersServer(cf string, service services.UsersService) (*UsersServer, error) {
	panic(wire.Build(testProviderSet))
}
