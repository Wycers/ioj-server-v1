// +build wireinject

package controllers

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/problems/repositories"
	"github.com/infinity-oj/server/internal/app/problems/services"
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

func CreateUsersController(cf string, sto repositories.ProblemRepository, client proto.FilesClient) (*ProblemController, error) {
	panic(wire.Build(testProviderSet))
}
