package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UsersRepository interface {
	Create(username, hash, salt, email string) (u *models.Account, err error)
	UpdateProfile(p *models.Profile) (err error)
	UpdateAccount(u *models.Account) (err error)
	QueryAccount(username string) (u *models.Account, err error)
	QueryProfile(uid uint64) (p *models.Profile, err error)
}

type MysqlUsersRepository struct {
	logger *zap.Logger
	db     *gorm.DB
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

	p := &models.Profile{
		UserId: u.ID,
		Email:  email,
	}
	if err = tx.Create(p).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, " create profile of user %s ", u.Username)
	}

	return u, tx.Commit().Error
}

func (s *MysqlUsersRepository) UpdateProfile(p *models.Profile) (err error) {
	// TODO: find a better way...
	err = s.db.Save(&p).Error
	return
}

func (s *MysqlUsersRepository) UpdateAccount(u *models.Account) (err error) {
	// TODO: find a better way...
	err = s.db.Save(&u).Error
	return
}

func (s *MysqlUsersRepository) QueryAccount(username string) (u *models.Account, err error) {
	u = new(models.Account)
	if err = s.db.Model(u).Where("username = ?", username).First(u).Error; err != nil {
		return nil, errors.Wrapf(err, "get profile error[Username = %s]", username)
	}
	return
}

func (s *MysqlUsersRepository) QueryProfile(uid uint64) (p *models.Profile, err error) {
	p = new(models.Profile)
	if err = s.db.Model(p).Where("UserId = ?", uid).First(p).Error; err != nil {
		return nil, errors.Wrapf(err, "get profile error[UserId = %d]", uid)
	}
	return
}

func NewMysqlUsersRepository(logger *zap.Logger, db *gorm.DB) UsersRepository {
	return &MysqlUsersRepository{
		logger: logger.With(zap.String("type", "UsersRepository")),
		db:     db,
	}
}
