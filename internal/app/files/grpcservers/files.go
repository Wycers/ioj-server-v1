package grpcservers

import (
	"context"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/files/services"
	"go.uber.org/zap"
)

type FilesServer struct {
	logger  *zap.Logger
	service services.FilesService
}

func (f FilesServer) CreateDirectory(context.Context, *proto.CreateDirectoryRequest) (*proto.CreateDirectoryResponse, error) {
	panic("implement me")
}

func (f FilesServer) CreateFile(ctx context.Context, req *proto.CreateFileRequest) (res *proto.CreateFileResponse, err error) {
	f.logger.Info("Create file " + req.GetFilePath() + " in " + req.GetFileSpace())

	if err := f.service.CreateFile(req.GetFileSpace(), req.GetFilePath(), req.GetData()); err != nil {
		res = &proto.CreateFileResponse{
			Status: 1,
		}
	} else {
		res = &proto.CreateFileResponse{
			Status: 0,
		}
	}
	return
}

func (f FilesServer) CreateFileSpace(ctx context.Context, req *proto.CreateFileSpaceRequest) (res *proto.CreateFileSpaceResponse, err error) {
	f.logger.Info("Create file space " + req.GetSpaceName())

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
