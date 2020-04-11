package services

import (
	"context"
	"fmt"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
)

type JudgementsService interface {
	Create(publicSpace, privateSpace, userSpace, testCase string) error
	Fetch() error
}

type DefaultJudgementsService struct {
	judgementsSrv proto.JudgementsClient
}

func (s *DefaultJudgementsService) Create(publicSpace, privateSpace, userSpace, testCase string) error {
	req := &proto.SubmitJudgementRequest{
		PublicSpace:  publicSpace,
		PrivateSpace: privateSpace,
		UserSpace:    userSpace,
		TestCase:     testCase,
	}

	res, err := s.judgementsSrv.SubmitJudgement(context.TODO(), req)
	fmt.Println(res.Status, res.Score)
	return err
}

func (s *DefaultJudgementsService) Fetch() error {
	panic("implement me")
}

func NewJudgementsService(judgementsSrv proto.JudgementsClient) JudgementsService {
	return &DefaultJudgementsService{
		judgementsSrv: judgementsSrv,
	}
}
