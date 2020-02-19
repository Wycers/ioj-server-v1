package services

import (
	"context"
	"fmt"
	"github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProductsService interface {
	Get(c context.Context, ID uint64) (*models.Account, error)
}

type DefaultProductsService struct {
	logger   *zap.Logger
	usersSrv proto.UsersClient
}

func NewProductService(logger *zap.Logger, usersSrv proto.UsersClient) ProductsService {
	return &DefaultProductsService{
		logger:   logger.With(zap.String("type", "DefaultProductsService")),
		usersSrv: usersSrv,
	}
}

func (s *DefaultProductsService) Get(c context.Context, productID uint64) (p *models.Account, err error) {
	// get detail
	req := &proto.RegisterRequest{
		Username: "wycer",
		Email:    "wycers@gmail.com",
		Password: "123",
	}

	pd, err := s.usersSrv.Register(c, req)
	if err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}

	fmt.Println(pd.User.Password)

	return nil, nil
}
