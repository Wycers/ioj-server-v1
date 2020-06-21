// +build wireinject

package services

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/submissions/repositories"
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

func CreateSubmissionsService(
	cf string,
	sto repositories.SubmissionRepository,
	problemsClient proto.ProblemsClient,
	filesClient proto.FilesClient,
	judgementsClient proto.JudgementsClient,
) (SubmissionsService, error) {
	panic(wire.Build(testProviderSet))
}
