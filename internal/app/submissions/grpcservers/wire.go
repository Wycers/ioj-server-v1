// +build wireinject

package grpcservers

import (
	"github.com/Infinity-OJ/Server/internal/app/submissions/services"
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

func CreateSubmissionsServer(cf string, service services.SubmissionsService) (*SubmissionService, error) {
	panic(wire.Build(testProviderSet))
}
