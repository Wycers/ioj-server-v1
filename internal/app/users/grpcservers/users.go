package grpcservers

import (
	"context"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/users/services"
	"github.com/infinity-oj/server/internal/pkg/jwt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersServer struct {
	logger  *zap.Logger
	service services.UsersService
}

func (s *UsersServer) mustEmbedUnimplementedUsersServer() {
	panic("implement me")
}

func (s *UsersServer) CreateUser(ctx context.Context, req *proto.RegisterRequest) (res *proto.RegisterResponse, err error) {
	if u, err := s.service.Create(req.Username, req.Password, req.Email); err != nil {
		return nil, errors.Wrapf(err, "CreateUser user failed")
	} else {
		res = &proto.RegisterResponse{
			User: &proto.User{
				Uid:      u.ID,
				Username: u.Username,
				Password: "",
			},
		}
	}
	return
}

func (s *UsersServer) CreateSession(ctx context.Context, req *proto.SigninRequest) (res *proto.SigninResponse, err error) {
	if isValid, err := s.service.Verify(req.Username, req.Password); err != nil {
		return nil, errors.Wrapf(err, "[CreateProblem Session] Query error")
	} else {
		if isValid {
			if token, err := jwt.GenerateToken(req.Username); err != nil {
				return nil, errors.Wrapf(err, "[CreateProblem Session] Generate Token failed")
			} else {
				res = &proto.SigninResponse{
					Authorized: true,
					Token:      token,
				}
			}
		} else {
			res = &proto.SigninResponse{
				Authorized: false,
				Token:      "",
			}
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
