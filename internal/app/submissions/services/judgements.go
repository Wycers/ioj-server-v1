package services

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
)

type JudgementsService interface {
	Fetch() error

	Create(class, property string, inputs [][]byte) error
}

type DefaultJudgementsService struct {
	judgementsSrv proto.JudgementsClient
}

//func (s *DefaultJudgementsService) Create(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error {
//	req := &proto.SubmitJudgementRequest{
//		SubmissionId: submissionId,
//		PublicSpace:  publicSpace,
//		PrivateSpace: privateSpace,
//		UserSpace:    userSpace,
//		TestCase:     testCase,
//	}
//
//	if res, err := s.judgementsSrv.SubmitJudgement(context.TODO(), req); err != nil {
//		return errors.Wrap(err, "judge error: submit judgement error")
//	} else {
//		fmt.Println(res.Status, res.Score)
//	}
//	return nil
//}
func (s *DefaultJudgementsService) Create(class, property string, inputs [][]byte) error {

	return nil
}

func (s *DefaultJudgementsService) Fetch() error {
	panic("implement me")
}

func NewJudgementsService(judgementsSrv proto.JudgementsClient) JudgementsService {
	return &DefaultJudgementsService{
		judgementsSrv: judgementsSrv,
	}
}
