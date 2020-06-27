package grpcservers

import (
	"github.com/google/wire"
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	us *UsersServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterUsersServer(s, us)
	}
}

var ProviderSet = wire.NewSet(NewUsersServer, CreateInitServersFn)
