// +build wireinject

package services

import (
	"github.com/google/wire"
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateJudgementsService(
	cf string,
	sto repositories.JudgementsRepository,
	filesClient proto.FilesClient,
	submissionsClient proto.SubmissionsClient,
) (JudgementsService, error) {
	panic(wire.Build(testProviderSet))
}
