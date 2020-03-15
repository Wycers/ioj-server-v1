// +build wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/server"
	"github.com/Infinity-OJ/Server/internal/app/server/controllers"
	"github.com/Infinity-OJ/Server/internal/app/server/grpcclients"
	"github.com/Infinity-OJ/Server/internal/app/server/services"
	"github.com/Infinity-OJ/Server/internal/pkg/app"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/consul"
	"github.com/Infinity-OJ/Server/internal/pkg/jaeger"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
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
