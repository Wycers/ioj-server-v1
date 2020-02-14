// +build wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/users"
	"github.com/Infinity-OJ/Server/internal/app/users/controllers"
	"github.com/Infinity-OJ/Server/internal/app/users/grpcservers"
	"github.com/Infinity-OJ/Server/internal/app/users/repositories"
	"github.com/Infinity-OJ/Server/internal/app/users/services"
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
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	users.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
