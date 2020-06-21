// +build wireinject

package grpcservers

import (
	"github.com/infinity-oj/server/internal/app/files/services"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateFilesServer(cf string, service services.FilesService) (*FilesServer, error) {
	panic(wire.Build(testProviderSet))
}
