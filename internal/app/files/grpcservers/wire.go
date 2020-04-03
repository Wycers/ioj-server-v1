// +build wireinject

package grpcservers

import (
	"github.com/Infinity-OJ/Server/internal/app/files/services"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
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
