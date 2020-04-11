package grpcclients

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/pkg/errors"
)

func NewFilesClient(client *grpc.Client) (proto.FilesClient, error) {
	conn, err := client.Dial("Files")
	if err != nil {
		return nil, errors.Wrap(err, "file client dial error")
	}
	c := proto.NewFilesClient(conn)

	return c, nil
}
