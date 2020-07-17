package services

import (
	proto "github.com/infinity-oj/api/protobuf-spec"
)

type SubmissionsService interface {
}

type DefaultSubmissionsService struct {
	submissionSrv proto.SubmissionsClient
}

func NewSubmissionsService(client proto.SubmissionsClient) SubmissionsService {
	return &DefaultSubmissionsService{
		submissionSrv: client,
	}
}
