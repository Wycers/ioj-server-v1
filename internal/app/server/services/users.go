package services

import (
	"context"
	"github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserService interface {
	Create(c context.Context, username, password, email string) (*models.Account, error)
	Update(c context.Context, ID uint64) error
}

type DefaultUserService struct {
	logger   *zap.Logger
	usersSrv proto.UsersClient
}

func (s *DefaultUserService) Create(c context.Context, username, password, email string) (*models.Account, error) {
	req := &proto.RegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}

	pd, err := s.usersSrv.Register(c, req)
	if err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}

	s.logger.Info("User Created!", zap.String("username", pd.User.Username))

	return &models.Account{
		Username: pd.User.Username,
	}, nil
}

func (d DefaultUserService) Update(c context.Context, ID uint64) error {
	panic("implement me")
}

func NewUserService(logger *zap.Logger, usersSrv proto.UsersClient) UserService {
	return &DefaultUserService{
		logger:   logger.With(zap.String("type", "DefaultProductsService")),
		usersSrv: usersSrv,
	}
}
