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

func (f FilesServer) CreateDirectory(ctx context.Context, req *proto.CreateDirectoryRequest) (res *proto.CreateDirectoryResponse, err error) {
	f.logger.Info("Create directory" + req.GetDirectory() + " in " + req.GetFileSpace())

	if err := f.service.CreateDirectory(req.GetFileSpace(), req.GetDirectory()); err != nil {
		f.logger.Error("Create directory error", zap.String("error:", err.Error()))
		res = &proto.CreateDirectoryResponse{
			Status: proto.Status_error,
		}
	} else {
		res = &proto.CreateDirectoryResponse{
			Status: proto.Status_success,
		}
	}
	return
}

func (f FilesServer) CreateFile(ctx context.Context, req *proto.CreateFileRequest) (res *proto.CreateFileResponse, err error) {
	f.logger.Info("Create file " + req.GetFilePath() + " in " + req.GetFileSpace())

	if err := f.service.CreateFile(req.GetFileSpace(), req.GetFilePath(), req.GetData()); err != nil {
		f.logger.Error("Create file error", zap.String("error:", err.Error()))
		res = &proto.CreateFileResponse{
			Status: proto.Status_error,
		}
	} else {
		res = &proto.CreateFileResponse{
			Status: proto.Status_success,
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
