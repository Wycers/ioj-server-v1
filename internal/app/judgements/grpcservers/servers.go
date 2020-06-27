package grpcservers

import (
	"github.com/google/wire"
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	js *JudgementsService,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterJudgementsServer(s, js)
	}
}

var ProviderSet = wire.NewSet(NewJudgementsServer, CreateInitServersFn)
