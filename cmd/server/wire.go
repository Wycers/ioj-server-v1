// +build wireinject

package main

import (
	"github.com/infinity-oj/server/internal/app/server"
	"github.com/infinity-oj/server/internal/app/server/controllers"
	"github.com/infinity-oj/server/internal/app/server/grpcclients"
	"github.com/infinity-oj/server/internal/app/server/services"
	"github.com/infinity-oj/server/internal/pkg/app"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/consul"
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
	http.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	server.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
