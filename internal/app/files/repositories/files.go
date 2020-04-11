package repositories

import (
	"path"

	"github.com/Infinity-OJ/Server/internal/pkg/files"
	"go.uber.org/zap"
)

type FilesRepository interface {
	CreateFileSpace(fileSpace string) error
	CreateDirectory(fileSpace, directory string) error
	CreateFile(fileSpace, fileName string, data []byte) error
	IsFileExists(fileSpace, fileName string) bool
	FetchFile(fileSpace, fileName string) ([]byte, error)
}

type FileManager struct {
	logger *zap.Logger
	fm     files.FileManager
}

func (m *FileManager) IsFileExists(fileSpace, fileName string) bool {
	filePath := path.Join(fileSpace, fileName)
	exist, err := m.fm.IsFileExists(filePath)
	if err != nil {
		return false
	}
	return exist
}

func (m *FileManager) FetchFile(fileSpace, fileName string) ([]byte, error) {
	panic("implement me")
}

func (m *FileManager) CreateDirectory(fileSpace, directory string) error {
	filePath := path.Join(fileSpace, directory)
	return m.fm.CreateDirectory(filePath)
}

func (m *FileManager) CreateFile(fileSpace, fileName string, data []byte) error {
	filePath := path.Join(fileSpace, fileName)
	return m.fm.CreateFile(filePath, data)
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
