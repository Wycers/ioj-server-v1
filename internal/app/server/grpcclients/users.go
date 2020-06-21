package grpcclients

import (
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewUsersClient(client *grpc.Client) (proto.UsersClient, error) {
	conn, err := client.Dial("Users")
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := proto.NewUsersClient(conn)

	return c, nil
}
