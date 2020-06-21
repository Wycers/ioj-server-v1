package grpcclients

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewJudgementsClient(client *grpc.Client) (proto.JudgementsClient, error) {
	conn, err := client.Dial("Judgements")
	if err != nil {
		return nil, errors.Wrap(err, "submission client dial error")
	}
	c := proto.NewJudgementsClient(conn)

	return c, nil
}
