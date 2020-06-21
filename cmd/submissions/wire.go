// +build wireinject

package main

import (
	"github.com/infinity-oj/server/internal/app/submissions"
	"github.com/infinity-oj/server/internal/app/submissions/controllers"
	"github.com/infinity-oj/server/internal/app/submissions/grpcclients"
	"github.com/infinity-oj/server/internal/app/submissions/grpcservers"
	"github.com/infinity-oj/server/internal/app/submissions/repositories"
	"github.com/infinity-oj/server/internal/app/submissions/services"
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
	database.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	repositories.ProviderSet,
	grpc.ProviderSet,
	grpcservers.ProviderSet,
	grpcclients.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	submissions.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
