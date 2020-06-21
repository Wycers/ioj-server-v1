// +build wireinject

package repositories

import (
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/files"
	"github.com/infinity-oj/server/internal/pkg/log"
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
