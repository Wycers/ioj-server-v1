package service

import (
	"context"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
)

type JudgementService interface {
	List() error
}

type DefaultJudgementService struct {
	judgementSrv proto.JudgementsClient
}

func (d DefaultJudgementService) List() error {

	req := &proto.ListRequest{}

	_, err := d.judgementSrv.ListJudgements(context.TODO(), req)
	return err
}

func NewJudgementService(client proto.JudgementsClient) JudgementService {
	return &DefaultJudgementService{
		judgementSrv: client,
	}
}
