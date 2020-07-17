// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
)

// Injectors from wire.go:

func CreateJudgementsService(cf string, sto repositories.JudgementsRepository, filesClient protobuf_spec.FilesClient, submissionsClient protobuf_spec.SubmissionsClient) (JudgementsService, error) {
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
	filesService := NewFilesService(filesClient)
	submissionsService := NewSubmissionsService(submissionsClient)
	judgementsService := NewJudgementsService(logger, sto, filesService, submissionsService)
	return judgementsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
