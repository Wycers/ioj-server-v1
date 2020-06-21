// +build wireinject

package grpcservers

import (
	"github.com/infinity-oj/server/internal/app/judgements/services"
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

func CreateJudgementsServer(cf string, service services.JudgementsService) (*JudgementsService, error) {
	panic(wire.Build(testProviderSet))
}
