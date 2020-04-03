package services

import (
	"github.com/Infinity-OJ/Server/internal/app/files/repositories"
	"go.uber.org/zap"
)

var specialKey = "imf1nlTy0j"

type FilesService interface {
	CreateFileSpace(fileSpace string) error
}

type DefaultFilesService struct {
	logger     *zap.Logger
	Repository repositories.FilesRepository
}

func (d DefaultFilesService) CreateFileSpace(fileSpace string) error {
	d.Repository.CreateFileSpace(fileSpace)
	return nil
}

func NewFileService(logger *zap.Logger, Repository repositories.FilesRepository) FilesService {
	return &DefaultFilesService{
		logger:     logger.With(zap.String("type", "DefaultFilesService")),
		Repository: Repository,
	}
}
