package controllers

import (
	"github.com/Infinity-OJ/Server/internal/app/submissions/services"
	"go.uber.org/zap"
)

type SubmissionController struct {
	logger  *zap.Logger
	service services.SubmissionsService
}

func NewSubmissionsController(logger *zap.Logger, s services.SubmissionsService) *SubmissionController {
	return &SubmissionController{
		logger:  logger,
		service: s,
	}
}
