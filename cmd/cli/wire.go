// +build wireinject

package main

import (
	"github.com/infinity-oj/server/internal/app/ctl"
	"github.com/infinity-oj/server/internal/app/ctl/commands/file"
	"github.com/infinity-oj/server/internal/app/ctl/commands/judgement"
	"github.com/infinity-oj/server/internal/app/ctl/commands/problem"
	"github.com/infinity-oj/server/internal/app/ctl/commands/submission"
	"github.com/infinity-oj/server/internal/app/ctl/commands/user"
	"github.com/infinity-oj/server/internal/app/ctl/grpcclients"
	"github.com/infinity-oj/server/internal/app/ctl/service"
	"github.com/infinity-oj/server/internal/pkg/app"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/consul"
	"github.com/infinity-oj/server/internal/pkg/jaeger"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"

	//"github.com/infinity-oj/server/internal/pkg/transports/http"
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
	judgement.ProviderSet,
	service.ProviderSet,
	consul.ProviderSet,
	ctl.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
)

func CreateApp(cf string) (*cli.App, error) {
	panic(wire.Build(providerSet))
}
