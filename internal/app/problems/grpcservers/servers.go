package grpcservers

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/google/wire"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ps *ProblemService,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterProblemsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewUsersServer, CreateInitServersFn)
