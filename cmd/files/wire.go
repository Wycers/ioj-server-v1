// +build wireinject

package main

import (
	"github.com/infinity-oj/server/internal/app/files"
	"github.com/infinity-oj/server/internal/app/files/controllers"
	"github.com/infinity-oj/server/internal/app/files/grpcservers"
	"github.com/infinity-oj/server/internal/app/files/repositories"
	"github.com/infinity-oj/server/internal/app/files/services"
	"github.com/infinity-oj/server/internal/pkg/app"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/consul"
	"github.com/infinity-oj/server/internal/pkg/database"
	file_manager "github.com/infinity-oj/server/internal/pkg/files"
	"github.com/infinity-oj/server/internal/pkg/jaeger"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/infinity-oj/server/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	files.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
	file_manager.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
