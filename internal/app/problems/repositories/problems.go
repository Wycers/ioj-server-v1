package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProblemRepository interface {
	Create(title string) (p *models.Problem, err error)
}

type MysqlUsersRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (m MysqlUsersRepository) Create(title string) (p *models.Problem, err error) {
	p = &models.Problem{Title: title}
	if err = m.db.Create(p).Error; err != nil {
		return nil, errors.Wrapf(err, "create problem with title: %s", title)
	}
	return p, nil
}

func NewMysqlUsersRepository(logger *zap.Logger, db *gorm.DB) ProblemRepository {
	return &MysqlUsersRepository{
		logger: logger.With(zap.String("type", "ProblemRepository")),
		db:     db,
	}
}
