package services

import (
	"github.com/Infinity-OJ/Server/internal/app/users/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersService interface {
	Get(ID uint64) (*models.Detail, error)
}

type DefaultUsersService struct {
	logger     *zap.Logger
	Repository repositories.UsersRepository
}

func NewDetailService(logger *zap.Logger, Repository repositories.UsersRepository) UsersService {
	return &DefaultUsersService{
		logger:     logger.With(zap.String("type", "DefaultUsersService")),
		Repository: Repository,
	}
}

func (s *DefaultUsersService) Get(ID uint64) (p *models.Detail, err error) {
	if p, err = s.Repository.Get(ID); err != nil {
		return nil, errors.Wrap(err, "detail service get detail error")
	}

	return
}
