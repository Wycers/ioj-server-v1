package grpcservers

import (
	"context"
	"github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/users/services"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersServer struct {
	logger  *zap.Logger
	service services.UsersService
}

func (s *UsersServer) Register(ctx context.Context, req *proto.RegisterRequest) (resp *proto.RegisterResponse, err error) {
	s.logger.Info(req.Username)
	s.logger.Info(req.Password)

	if u, err := s.service.Create(req.Username, req.Password, req.Email); err != nil {
		return nil, errors.Wrapf(err, "Create user failed")
	} else {
		resp = &proto.RegisterResponse{
			User: &proto.User{
				Uid:      u.ID,
				Username: u.Username,
				Password: "",
			},
		}
	}
	return
}

func NewUsersServer(logger *zap.Logger, ps services.UsersService) (*UsersServer, error) {
	return &UsersServer{
		logger:  logger,
		service: ps,
	}, nil
}
