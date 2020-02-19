package grpcclients

import (
	"github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
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
