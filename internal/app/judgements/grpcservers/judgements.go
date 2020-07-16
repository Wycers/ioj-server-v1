package grpcservers

import (
	"context"
	"errors"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"go.uber.org/zap"

	"github.com/infinity-oj/server/internal/app/judgements/services"
)

type JudgementsService struct {
	logger  *zap.Logger
	service services.JudgementsService
}

func (s *JudgementsService) ListJudgements(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error) {
	s.service.List()
	return &proto.ListResponse{}, nil
}

func (s *JudgementsService) CreateJudgement(ctx context.Context, request *proto.CreateJudgementRequest) (*proto.CreateJudgementResponse, error) {
	tp := request.Type

	arguments := make(map[string]string)
	for _, v := range request.Arguments {
		if _, ok := arguments[v.Key]; ok {
			return nil, errors.New("duplicate key")
		}
		arguments[v.Key] = v.Value
	}

	var inputs [][]byte
	for _, v := range request.Slots {
		input := v.Value
		inputs = append(inputs, input)
	}

	_, err := s.service.Create(tp, arguments, inputs)
	if err != nil {
		return nil, err
	}

	response := &proto.CreateJudgementResponse{}

	return response, nil
}

func (s *JudgementsService) PullJudgement(ctx context.Context, request *proto.PullJudgementRequest) (*proto.PullJudgementResponse, error) {
	taskType := request.GetType()

	token, judgement := s.service.PullJudgement(taskType)

	if judgement == nil {
		return nil, nil
	}

	var arguments []*proto.Argument
	for k, v := range judgement.Properties {
		argument := &proto.Argument{
			Key:   k,
			Value: v,
		}
		arguments = append(arguments, argument)
	}

	var slots []*proto.Slot
	for k, v := range judgement.Inputs {
		slot := &proto.Slot{
			Id:    uint32(k),
			Value: v,
		}
		slots = append(slots, slot)
	}

	rsp := &proto.PullJudgementResponse{
		Token:     token,
		Arguments: arguments,
		Slots:     slots,
	}

	return rsp, nil
}

func (s *JudgementsService) PushJudgement(ctx context.Context, request *proto.PushJudgementRequest) (*proto.PushJudgementResponse, error) {
	token := request.Token

	var outputs [][]byte
	for _, v := range request.Slots {
		output := v.Value
		outputs = append(outputs, output)
	}

	err := s.service.PushJudgement(token, outputs)
	if err != nil {
		return nil, err
	}

	response := &proto.PushJudgementResponse{
		Status: proto.Status_success,
	}
	return response, nil
}

func NewJudgementsServer(logger *zap.Logger, js services.JudgementsService) (*JudgementsService, error) {
	return &JudgementsService{
		logger:  logger,
		service: js,
	}, nil
}
