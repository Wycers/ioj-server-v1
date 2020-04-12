package grpcclients

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewJudgementsClient(client *grpc.Client) (proto.JudgementsClient, error) {
	conn, err := client.Dial("Judgements")
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := proto.NewJudgementsClient(conn)

	return c, nil
}
