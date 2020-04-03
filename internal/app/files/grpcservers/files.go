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

func (s *FilesServer) CreateFileSpace(context.Context, *proto.CreateFileSpaceRequest) (*proto.CreateFileSpaceResponse, error) {
	panic("implement me")
}

func NewFilesServer(logger *zap.Logger, fs services.FilesService) (*FilesServer, error) {
	return &FilesServer{
		logger:  logger,
		service: fs,
	}, nil
}
