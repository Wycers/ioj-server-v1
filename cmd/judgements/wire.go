// +build wireinject

package main

import (
	"github.com/infinity-oj/server/internal/app/judgements"
	"github.com/infinity-oj/server/internal/app/judgements/controllers"
	"github.com/infinity-oj/server/internal/app/judgements/grpcclients"
	"github.com/infinity-oj/server/internal/app/judgements/grpcservers"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"github.com/infinity-oj/server/internal/app/judgements/services"
	"github.com/infinity-oj/server/internal/pkg/app"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/consul"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/jaeger"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/infinity-oj/server/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	database.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	grpcservers.ProviderSet,
	grpcclients.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	judgements.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
