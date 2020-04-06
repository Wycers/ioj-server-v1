package files

import (
	"errors"
	"io/ioutil"
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

func (m *LocalFileManager) CreateFile(fileName string, bytes []byte) (err error) {
	filePath, err := GetFileAbsPath(m.base, fileName)
	if err != nil {
		return
	}
	if exist, err := m.isFileExists(filePath); err != nil {
		return err
	} else {
		if exist {
			return errors.New("file or directory exists")
		} else {
			err = ioutil.WriteFile(filePath, bytes, os.FileMode(0644))
		}
	}
	return
}

func (m *LocalFileManager) CreateDirectory(fileName string) (err error) {
	filePath, err := GetFileAbsPath(m.base, fileName)
	if err != nil {
		return
	}
	if exist, err := m.isFileExists(filePath); err != nil {
		return err
	} else {
		if exist {
			return errors.New("file or directory exists")
		} else {
			err = os.MkdirAll(filePath, os.FileMode(0644))
		}
	}
	return
}

func (m *LocalFileManager) isFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
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
