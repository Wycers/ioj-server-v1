// +build wireinject

package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
	"github.com/Infinity-OJ/Server/internal/pkg/files"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	files.ProviderSet,
	ProviderSet,
)

func CreateFileRepository(f string) (FilesRepository, error) {
	panic(wire.Build(testProviderSet))
}
