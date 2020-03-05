package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersRepository interface {
	Query(ID uint64) (p *models.Detail, err error)
	Create(username, hash, salt, email string) (u *models.Account, err error)
}

type MysqlUsersRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlUsersRepository(logger *zap.Logger, db *gorm.DB) UsersRepository {
	return &MysqlUsersRepository{
		logger: logger.With(zap.String("type", "UsersRepository")),
		db:     db,
	}
}

func (s *MysqlUsersRepository) Query(ID uint64) (p *models.Detail, err error) {
	p = new(models.Detail)
	if err = s.db.Model(p).Where("id = ?", ID).First(p).Error; err != nil {
		return nil, errors.Wrapf(err, "get product error[id=%d]", ID)
	}
	return
}

func (s *MysqlUsersRepository) Create(username, hash, salt, email string) (u *models.Account, err error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	u = &models.Account{
		Username: username,
		Hash:     hash,
		Salt:     salt,
	}
	if err = tx.Create(u).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, " create user with username: %s", u.Username)
	}

	p := &models.Profile{Email: email}
	if err = tx.Create(p).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, " create profile of user %s ", u.Username)
	}

	return u, tx.Commit().Error
}
