// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server/internal/app/users/services"
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

func CreateUsersServer(cf string, service services.UsersService) (*UsersServer, error) {
	panic(wire.Build(testProviderSet))
}
