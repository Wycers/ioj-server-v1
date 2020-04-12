package services

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
)

type FilesService interface {
	FetchFile(fileSpace, fileName string) (data []byte, err error)
}

type DefaultFilesService struct {
	fileSrv proto.FilesClient
}

func (d DefaultFilesService) FetchFile(fileSpace, fileName string) (data []byte, err error) {
	req := &proto.FetchFileRequest{
		SpaceName: fileSpace,
		FilePath:  fileName,
	}

	res, err := d.fileSrv.FetchFile(context.TODO(), req)
	if err != nil {
		return nil, errors.Wrap(err, "fetch meta file error")
	}
	return res.GetFile().GetData(), nil
}

func NewFilesService(client proto.FilesClient) FilesService {
	return &DefaultFilesService{
		fileSrv: client,
	}
}
