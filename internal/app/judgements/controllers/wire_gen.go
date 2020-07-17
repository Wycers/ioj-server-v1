// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"github.com/infinity-oj/server/internal/app/judgements/services"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
)

// Injectors from wire.go:

func CreateJudgementsController(cf string, sto repositories.JudgementsRepository, fileClient protobuf_spec.FilesClient, submissionClient protobuf_spec.SubmissionsClient) (*JudgementController, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	filesService := services.NewFilesService(fileClient)
	submissionsService := services.NewSubmissionsService(submissionClient)
	judgementsService := services.NewJudgementsService(logger, sto, filesService, submissionsService)
	judgementController := NewJudgementsController(logger, judgementsService)
	return judgementController, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, ProviderSet)
