// +build wireinject

package services

import (
	"github.com/Infinity-OJ/Server/internal/app/files/repositories"
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

func CreateFilesService(cf string, sto repositories.FilesRepository) (FilesService, error) {
	panic(wire.Build(testProviderSet))
}
