package grpcservers

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/google/wire"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ss *SubmissionService,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterSubmissionsServer(s, ss)
	}
}

var ProviderSet = wire.NewSet(NewUsersServer, CreateInitServersFn)
