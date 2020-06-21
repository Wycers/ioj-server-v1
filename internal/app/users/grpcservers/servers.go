package grpcservers

import (
	"github.com/google/wire"
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ps *UsersServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterUsersServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewUsersServer, CreateInitServersFn)
