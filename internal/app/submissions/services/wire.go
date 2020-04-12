// +build wireinject

package services

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/submissions/repositories"
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

func CreateSubmissionsService(
	cf string,
	sto repositories.SubmissionRepository,
	problemsClient proto.ProblemsClient,
	filesClient proto.FilesClient,
	judgementsClient proto.JudgementsClient,
) (SubmissionsService, error) {
	panic(wire.Build(testProviderSet))
}
