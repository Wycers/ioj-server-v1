package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProblemRepository interface {
	CreateProblem(title, locale, publicSpace, privateSpace string) (p *models.Page, err error)
	FindProblemById(problemId string) (p *models.Problem, err error)
}

type MysqlUsersRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (m MysqlUsersRepository) FindProblemById(problemId string) (p *models.Problem, err error) {
	p = &models.Problem{}
	if err = m.db.Where("problem_id = ?", problemId).First(&p).Error; err != nil {
		p = nil
	}
	return
}

func (m MysqlUsersRepository) CreateProblem(title, locale, publicSpace, privateSpace string) (page *models.Page, err error) {
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
		Locale:       locale,
		ProblemId:    random.RandStringRunes(8),
		PublicSpace:  publicSpace,
		PrivateSpace: privateSpace,
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
