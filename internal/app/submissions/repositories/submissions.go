package repositories

import (
	"github.com/infinity-oj/server/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SubmissionRepository interface {
	Create(submitterID uint64, problemId, userSpace string) (s *models.Submission, err error)
	Update(s *models.Submission) error
}

type MysqlSubmissionsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (m MysqlSubmissionsRepository) Create(submitterId uint64, problemId, userSpace string) (s *models.Submission, err error) {
	s = &models.Submission{
		SubmitterId: submitterId,
		ProblemID:   problemId,
		Judgements:  nil,
		UserSpace:   userSpace,
		Status:      models.Pending,
	}
	if err = m.db.Create(s).Error; err != nil {
		return nil, errors.Wrapf(err,
			" create submission with username: %d, problemID: %s, userSpace: %s",
			submitterId, problemId, userSpace,
		)
	}
	return
}

func (m MysqlSubmissionsRepository) Update(s *models.Submission) (err error) {
	err = m.db.Save(s).Error
	return
}

func NewMysqlSubmissionsRepository(logger *zap.Logger, db *gorm.DB) SubmissionRepository {
	return &MysqlSubmissionsRepository{
		logger: logger.With(zap.String("type", "SubmissionRepository")),
		db:     db,
	}
}
