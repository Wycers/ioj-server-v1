// +build wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/judgements"
	"github.com/Infinity-OJ/Server/internal/app/judgements/controllers"
	"github.com/Infinity-OJ/Server/internal/app/judgements/grpcclients"
	"github.com/Infinity-OJ/Server/internal/app/judgements/grpcservers"
	"github.com/Infinity-OJ/Server/internal/app/judgements/repositories"
	"github.com/Infinity-OJ/Server/internal/app/judgements/services"
	"github.com/Infinity-OJ/Server/internal/pkg/app"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/consul"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
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
