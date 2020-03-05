package services

import (
	"github.com/Infinity-OJ/Server/internal/app/users/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/crypto"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var specialKey = "imf1nlTy0j"

type UsersService interface {
	Get(ID uint64) (*models.Detail, error)
	Create(username, password, email string) (*models.Account, error)
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
	if p, err = s.Repository.Query(ID); err != nil {
		return nil, errors.Wrap(err, "detail service get detail error")
	}

	return
}

func (s *DefaultUsersService) Create(username, password, email string) (u *models.Account, err error) {
	salt := random.RandStringRunes(64)
	hash := crypto.Sha1(salt + password + specialKey)
	if u, err = s.Repository.Create(username, hash, salt, email); err != nil {
		return nil, err
	}
	return
}
