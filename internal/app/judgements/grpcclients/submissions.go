package grpcclients

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewSubmissionsClient(client *grpc.Client) (proto.SubmissionsClient, error) {
	conn, err := client.Dial("Submissions")
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := proto.NewSubmissionsClient(conn)

	return c, nil
}
