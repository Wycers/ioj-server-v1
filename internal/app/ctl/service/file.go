package service

import (
	"context"
	"fmt"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/pkg/errors"
)

type FileService interface {
	Create(fileSpace string) error
}

type DefaultFileService struct {
	fileSrv proto.FilesClient
}

func (d *DefaultFileService) Create(fileSpace string) error {
	req := &proto.CreateFileSpaceRequest{
		SpaceName: fileSpace,
	}

	fs, err := d.fileSrv.CreateFileSpace(context.TODO(), req)
	if err != nil {
		return errors.Wrap(err, "get rating error")
	}
	fmt.Println(fs.Status)
	return nil
}

func NewFileService(fileSrv proto.FilesClient) FileService {
	return &DefaultFileService{
		fileSrv: fileSrv,
	}
}
