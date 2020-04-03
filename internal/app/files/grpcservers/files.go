package grpcservers

import (
	"context"
	"fmt"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/files/services"
	"go.uber.org/zap"
)

type FilesServer struct {
	logger  *zap.Logger
	service services.FilesService
}

func (f FilesServer) CreateFileSpace(ctx context.Context, req *proto.CreateFileSpaceRequest) (res *proto.CreateFileSpaceResponse, err error) {
	fmt.Println(req.SpaceName)
	if err := f.service.CreateFileSpace(req.SpaceName); err != nil {
		res = &proto.CreateFileSpaceResponse{
			Status: 1,
		}
	} else {
		res = &proto.CreateFileSpaceResponse{
			Status: 0,
		}
	}
	return
}

func NewFilesServer(logger *zap.Logger, fs services.FilesService) (*FilesServer, error) {
	return &FilesServer{
		logger:  logger,
		service: fs,
	}, nil
}
