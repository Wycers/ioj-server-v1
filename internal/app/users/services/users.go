package services

import (
	"github.com/Infinity-OJ/Server/internal/app/users/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/crypto"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"
	"go.uber.org/zap"
)

var specialKey = "imf1nlTy0j"

type UsersService interface {
	//Get(ID uint64) (*models.Detail, error)
	Create(username, password, email string) (*models.User, error)
	Verify(username, password string) (valid bool, err error)
}

type DefaultUsersService struct {
	logger     *zap.Logger
	Repository repositories.UsersRepository
}

func (s *DefaultUsersService) Create(username, password, email string) (u *models.User, err error) {
	salt := random.RandStringRunes(64)
	hash := crypto.Sha1(salt + password + specialKey)
	if u, err = s.Repository.Create(username, hash, salt, email); err != nil {
		return nil, err
	}
	return
}

func (s *DefaultUsersService) Verify(username, password string) (valid bool, err error) {
	u := new(models.User)
	if u, err = s.Repository.QueryAccount(username); err != nil {
		return false, err
	}
	hash := crypto.Sha1(u.Salt + password + specialKey)

	return hash == u.Hash, nil
}

func NewDetailService(logger *zap.Logger, Repository repositories.UsersRepository) UsersService {
	return &DefaultUsersService{
		logger:     logger.With(zap.String("type", "DefaultUsersService")),
		Repository: Repository,
	}
}
