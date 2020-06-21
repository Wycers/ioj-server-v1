// +build wireinject

package services

import (
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	ProviderSet,
)

func CreateProductsService(cf string,
	usersClients proto.UsersClient) (ProductsService, error) {
	panic(wire.Build(testProviderSet))
}
