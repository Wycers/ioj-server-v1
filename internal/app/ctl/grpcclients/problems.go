package grpcclients

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewProblemsClient(client *grpc.Client) (proto.ProblemsClient, error) {
	conn, err := client.Dial("Problems")
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := proto.NewProblemsClient(conn)

	return c, nil
}
