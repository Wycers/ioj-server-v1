package services

import (
	"context"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(c context.Context, username, password, email string) (*models.User, error)
	Update(c context.Context, ID uint64) error
	CreateSession(c context.Context, username, password string) (token string, err error)
}

type DefaultUserService struct {
	logger   *zap.Logger
	usersSrv proto.UsersClient
}

func (s *DefaultUserService) CreateSession(c context.Context, username, password string) (token string, err error) {
	req := &proto.SigninRequest{
		Username: username,
		Password: password,
	}

	resp, err := s.usersSrv.CreateSession(c, req)
	if err != nil {
		return "", errors.Wrap(err, "Sign in failed.")
	}

	s.logger.Info("User Sign in!", zap.String("username", username))

	return resp.Token, nil
}

func (s *DefaultUserService) CreateUser(c context.Context, username, password, email string) (*models.User, error) {
	req := &proto.RegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}

	pd, err := s.usersSrv.CreateUser(c, req)
	if err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}

	s.logger.Info("User Created!", zap.String("username", pd.User.Username))

	return &models.User{
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
