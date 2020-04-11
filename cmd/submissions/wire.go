// +build wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/submissions"
	"github.com/Infinity-OJ/Server/internal/app/submissions/controllers"
	"github.com/Infinity-OJ/Server/internal/app/submissions/grpcclients"
	"github.com/Infinity-OJ/Server/internal/app/submissions/grpcservers"
	"github.com/Infinity-OJ/Server/internal/app/submissions/repositories"
	"github.com/Infinity-OJ/Server/internal/app/submissions/services"
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
