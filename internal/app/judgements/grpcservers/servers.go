package grpcservers

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/google/wire"
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
