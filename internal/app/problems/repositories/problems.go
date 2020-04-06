package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProblemRepository interface {
	Create(title, locale string) (p *models.Page, err error)
}

type MysqlUsersRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (m MysqlUsersRepository) Create(title, locale string) (page *models.Page, err error) {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	problem := &models.Problem{
		Locale:    locale,
		ProblemID: random.RandStringRunes(8),
		FileSpace: "",
	}
	if err = tx.Create(problem).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, " create problem with locale %s failed", locale)
	}

	page = &models.Page{Title: title}
	if err = m.db.Create(page).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, "create problem page with title: %s", title)
	}

	return page, tx.Commit().Error
}

func NewMysqlUsersRepository(logger *zap.Logger, db *gorm.DB) ProblemRepository {
	return &MysqlUsersRepository{
		logger: logger.With(zap.String("type", "ProblemRepository")),
		db:     db,
	}
}
