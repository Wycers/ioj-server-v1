package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersRepository interface {
	Get(ID uint64) (p *models.Detail, err error)
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

func (s *MysqlUsersRepository) Get(ID uint64) (p *models.Detail, err error) {
	p = new(models.Detail)
	if err = s.db.Model(p).Where("id = ?", ID).First(p).Error; err != nil {
		return nil, errors.Wrapf(err, "get product error[id=%d]", ID)
	}
	return
}
