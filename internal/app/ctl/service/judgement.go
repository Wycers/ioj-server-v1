package service

import (
	"context"

	proto "github.com/infinity-oj/api/protobuf-spec"
)

type JudgementService interface {
	List() error
	Create(submissionId uint64, publicSpace, privateSpace, userSpace string) error
}

type DefaultJudgementService struct {
	judgementSrv proto.JudgementsClient
}

func (d DefaultJudgementService) List() error {

	req := &proto.ListRequest{}

	_, err := d.judgementSrv.ListJudgements(context.TODO(), req)
	return err
}

func (d DefaultJudgementService) Create(submissionId uint64, publicSpace, privateSpace, userSpace string) error {
	// req := &proto.Create{
	// 	SubmissionId: submissionId,
	// 	PublicSpace:  publicSpace,
	// 	PrivateSpace: privateSpace,
	// 	UserSpace:    userSpace,
	// 	TestCase:     "",
	// }
	//
	// if res, err := d.judgementSrv.SubmitJudgement(context.TODO(), req); err != nil {
	// 	return err
	// } else {
	// 	fmt.Println(res.Status, res.Score)
	// }
	return nil
}

func NewJudgementService(client proto.JudgementsClient) JudgementService {
	return &DefaultJudgementService{
		judgementSrv: client,
	}
}
