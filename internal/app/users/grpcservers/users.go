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

	if u, err := s.service.Create(req.Username, req.Password, ""); err != nil {
		return nil, errors.Wrapf(err, "Create user failed")
	} else {
		resp = &proto.RegisterResponse{
			User: &proto.User{
				Uid:      0,
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

func (s *UsersServer) Get(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	//p, err := s.service.Get(req.Id)
	//if err != nil {
	//	return nil, errors.Wrap(err, "users grpc service get detail error")
	//}
	//ct, err := ptypes.TimestampProto(p.CreatedTime)
	//if err != nil {
	//	return nil, errors.Wrap(err, "convert create time error")
	//}
	//
	//resp := &proto.User{
	//	Id:          uint64(p.ID),
	//	Name:        p.Name,
	//	Price:       p.Price,
	//	CreatedTime: ct,
	//}
	s.logger.Info(req.Username)
	s.logger.Info(req.Password)
	resp := &proto.RegisterResponse{
		User: &proto.User{
			Uid:      123,
			Username: "123",
			Password: "123",
		},
	}
	return resp, nil
}
