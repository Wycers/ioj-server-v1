package services

import (
	"github.com/Infinity-OJ/Server/internal/app/files/repositories"
	"go.uber.org/zap"
)

var specialKey = "imf1nlTy0j"

type FilesService interface {
	CreateFileSpace(fileSpace string) error
	CreateDirectory(fileSpace, directory string) error
	CreateFile(fileSpace, fileName string, data []byte) error
}

type DefaultFilesService struct {
	logger     *zap.Logger
	Repository repositories.FilesRepository
}

func (d DefaultFilesService) CreateDirectory(fileSpace, directory string) error {
	panic("implement me")
}

func (d DefaultFilesService) CreateFile(fileSpace, fileName string, data []byte) error {
	return d.Repository.CreateFile(fileSpace, fileName, data)
}

func (d DefaultFilesService) CreateFileSpace(fileSpace string) error {
	return d.Repository.CreateFileSpace(fileSpace)
}

func NewFileService(logger *zap.Logger, Repository repositories.FilesRepository) FilesService {
	return &DefaultFilesService{
		logger:     logger.With(zap.String("type", "DefaultFilesService")),
		Repository: Repository,
	}
}
