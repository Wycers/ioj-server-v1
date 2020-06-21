package grpcservers

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/google/wire"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	fs *FilesServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterFilesServer(s, fs)
	}
}

var ProviderSet = wire.NewSet(NewFilesServer, CreateInitServersFn)
