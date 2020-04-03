// +build wireinject

package controllers

import (
	"github.com/Infinity-OJ/Server/internal/app/files/repositories"
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
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)

func CreateFilesController(cf string, sto repositories.FilesRepository) (*FilesController, error) {
	panic(wire.Build(testProviderSet))
}
