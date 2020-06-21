package grpcclients

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewFilesClient(client *grpc.Client) (proto.FilesClient, error) {
	conn, err := client.Dial("Files")
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := proto.NewFilesClient(conn)

	return c, nil
}
