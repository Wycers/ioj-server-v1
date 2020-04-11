// +build wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/ctl"
	"github.com/Infinity-OJ/Server/internal/app/ctl/commands/file"
	"github.com/Infinity-OJ/Server/internal/app/ctl/commands/problem"
	"github.com/Infinity-OJ/Server/internal/app/ctl/commands/submission"
	"github.com/Infinity-OJ/Server/internal/app/ctl/commands/user"
	"github.com/Infinity-OJ/Server/internal/app/ctl/grpcclients"
	"github.com/Infinity-OJ/Server/internal/app/ctl/service"
	"github.com/Infinity-OJ/Server/internal/pkg/app"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/consul"
	"github.com/Infinity-OJ/Server/internal/pkg/jaeger"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"

	//"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var providerSet = wire.NewSet(
	jaeger.ProviderSet,
	app.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	user.ProviderSet,
	file.ProviderSet,
	problem.ProviderSet,
	submission.ProviderSet,
	service.ProviderSet,
	consul.ProviderSet,
	ctl.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
)

func CreateApp(cf string) (*cli.App, error) {
	panic(wire.Build(providerSet))
}
