// +build wireinject

package controllers

import (
	"github.com/google/wire"
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"github.com/infinity-oj/server/internal/app/judgements/services"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	ProviderSet,
)

func CreateJudgementsController(
	cf string,
	sto repositories.JudgementsRepository,
	fileClient proto.FilesClient,
	submissionClient proto.SubmissionsClient,
) (*JudgementController, error) {
	panic(wire.Build(testProviderSet))
}
