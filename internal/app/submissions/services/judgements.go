package services

import (
	"context"
	"fmt"

	proto "github.com/infinity-oj/api/protobuf-spec"
)

type JudgementsService interface {
	Create(ctx context.Context, tp string, properties map[string]string, inputs [][]byte) (string, error)
}

type DefaultJudgementsService struct {
	judgementsSrv proto.JudgementsClient
}

func (s *DefaultJudgementsService) Create(
	ctx context.Context, tp string, properties map[string]string, inputs [][]byte,
) (string, error) {

	var arguments []*proto.Argument
	for k, v := range properties {
		argument := &proto.Argument{
			Key:   k,
			Value: v,
		}
		arguments = append(arguments, argument)
	}

	var slots []*proto.Slot
	for k, v := range inputs {
		slot := &proto.Slot{
			Id:    uint32(k),
			Value: v,
		}
		slots = append(slots, slot)
	}

	request := &proto.CreateJudgementRequest{
		SubmissionId: 0,
		Type:         tp,
		Arguments:    arguments,
		Slots:        slots,
	}
	fmt.Println(tp)

	resp, err := s.judgementsSrv.CreateJudgement(ctx, request)
	if err != nil {
		return "", err
	}

	return resp.JudgementId, nil
}

func NewJudgementsService(judgementsSrv proto.JudgementsClient) JudgementsService {
	return &DefaultJudgementsService{
		judgementsSrv: judgementsSrv,
	}
}
