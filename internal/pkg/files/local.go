package files

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type LocalFileManager struct {
	base string
}

func (m *LocalFileManager) GetBase() string {
	return m.base
}

func (m *LocalFileManager) SetBase(base string) {
	m.base = base
}

func GetFileAbsPath(base, fileName string) (fileAbsPath string, err error) {
	fileAbsPath = ""

	spacePath := path.Join(base)
	spaceAbsPath, err := filepath.Abs(spacePath)
	if err != nil {
		return
	}

	filePath := path.Join(base, fileName)
	fileAbsPath, err = filepath.Abs(filePath)
	if err != nil {
		return
	}

	if !strings.HasPrefix(fileAbsPath, spaceAbsPath) {
		return "", errors.New("escape from base path")
	}
	return
}

func (m *LocalFileManager) CreateFile(fileName string) error {
	panic("implement me")
}

func (m *LocalFileManager) CreateDirectory(fileName string) (err error) {
	filePath, err := GetFileAbsPath(m.base, fileName)
	if err != nil {
		return
	}
	fmt.Println(filePath)
	if exist, err := m.isFileExists(fileName); err != nil {
		return err
	} else {
		if exist {
			return errors.New("file or directory exists")
		} else {
			err = os.MkdirAll(filePath, os.ModePerm)
		}
	}
	return
}

func (m *LocalFileManager) isFileExists(fileName string) (bool, error) {
	path, err := GetFileAbsPath(m.base, fileName)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func (m *LocalFileManager) isDirectoryExists(fileName string) (bool, error) {
	panic("implement me")
}
