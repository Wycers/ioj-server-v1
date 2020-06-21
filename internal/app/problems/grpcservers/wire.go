// +build wireinject

package grpcservers

import (
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
	ProviderSet,
)

func CreateUsersServer(cf string, service services.ProblemsService) (*ProblemService, error) {
	panic(wire.Build(testProviderSet))
}
