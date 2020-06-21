// +build wireinject

package controllers

import (
	"github.com/infinity-oj/server/internal/app/files/repositories"
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
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)

func CreateFilesController(cf string, sto repositories.FilesRepository) (*FilesController, error) {
	panic(wire.Build(testProviderSet))
}
