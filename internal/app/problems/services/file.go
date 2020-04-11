package services

import (
	"context"
	"fmt"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/pkg/errors"
)

type FileService interface {
	CreateFileSpace(fileSpace string) error
	CreateDirectory(fileSpace, directory string) error
	CreateFile(fileSpace, fileName string, data []byte) error
}

type DefaultFileService struct {
	fileSrv proto.FilesClient
}

func (d *DefaultFileService) CreateDirectory(fileSpace, directory string) error {
	req := &proto.CreateDirectoryRequest{
		FileSpace: fileSpace,
		Directory: directory,
	}

	fs, err := d.fileSrv.CreateDirectory(context.TODO(), req)
	if err != nil {
		return errors.Wrap(err, "create directory error")
	}
	fmt.Println(fs.Status)
	return nil
}

func (d *DefaultFileService) CreateFile(fileSpace, fileName string, data []byte) error {
	req := &proto.CreateFileRequest{
		FileSpace: fileSpace,
		FilePath:  fileName,
		Data:      data,
	}

	fs, err := d.fileSrv.CreateFile(context.TODO(), req)
	if err != nil {
		return errors.Wrap(err, "Create file error")
	}
	fmt.Println(fs.Status)
	return nil
}

func (d *DefaultFileService) CreateFileSpace(fileSpace string) error {
	req := &proto.CreateFileSpaceRequest{
		SpaceName: fileSpace,
	}

	fs, err := d.fileSrv.CreateFileSpace(context.TODO(), req)
	if err != nil {
		return errors.Wrap(err, "create file space error")
	}
	fmt.Println(fs.Status)
	return nil
}

func NewFileService(fileSrv proto.FilesClient) FileService {
	return &DefaultFileService{
		fileSrv: fileSrv,
	}
}
