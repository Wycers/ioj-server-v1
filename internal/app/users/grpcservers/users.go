package grpcservers

import (
	"context"
	"github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/users/services"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersServer struct {
	logger  *zap.Logger
	service services.UsersService
}

func NewUsersServer(logger *zap.Logger, ps services.UsersService) (*UsersServer, error) {
	return &UsersServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *UsersServer) Get(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	p, err := s.service.Get(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "users grpc service get detail error")
	}
	ct, err := ptypes.TimestampProto(p.CreatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "convert create time error")
	}

	resp := &proto.User{
		Id:          uint64(p.ID),
		Name:        p.Name,
		Price:       p.Price,
		CreatedTime: ct,
	}

	return resp, nil
}
