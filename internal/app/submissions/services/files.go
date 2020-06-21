package services

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/infinity-oj/api/protobuf-spec"
)

type FilesService interface {
	FetchMetaFile(privateSpace string) (data []byte, err error)
}

type DefaultFilesService struct {
	fileSrv proto.FilesClient
}

func (d DefaultFilesService) FetchMetaFile(privateSpace string) (data []byte, err error) {
	req := &proto.FetchFileRequest{
		SpaceName: privateSpace,
		FilePath:  "meta.yaml",
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
