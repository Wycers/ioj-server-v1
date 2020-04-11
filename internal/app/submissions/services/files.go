package services

import (
	"context"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
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
		FilePath:  "meta.yml",
	}

	res, err := d.fileSrv.FetchFile(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return res.GetFile().GetData(), nil
}

func NewFilesService(client proto.FilesClient) FilesService {
	return &DefaultFilesService{
		fileSrv: client,
	}
}
