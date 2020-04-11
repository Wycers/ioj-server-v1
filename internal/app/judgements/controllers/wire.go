// +build wireinject

package controllers

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/judgements/repositories"
	"github.com/Infinity-OJ/Server/internal/app/judgements/services"
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

func CreateJudgementsController(cf string, sto repositories.JudgementsRepository, client proto.FilesClient) (*JudgementController, error) {
	panic(wire.Build(testProviderSet))
}
