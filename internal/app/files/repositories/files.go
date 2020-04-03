package repositories

import (
	"github.com/Infinity-OJ/Server/internal/pkg/files"
	"go.uber.org/zap"
)

type FilesRepository interface {
	CreateFileSpace(fileSpace string) error
}

type FileManager struct {
	logger *zap.Logger
	fm     files.FileManager
}

func (m *FileManager) CreateFileSpace(fileSpace string) error {
	err := m.fm.CreateDirectory(fileSpace)
	if err != nil {
		return err
	}
	return nil
}

func NewFileManager(logger *zap.Logger, fm files.FileManager) FilesRepository {
	return &FileManager{
		logger: logger.With(zap.String("type", "FilesRepository")),
		fm:     fm,
	}
}
