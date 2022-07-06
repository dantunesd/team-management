package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"team-management/members/api"
	"team-management/members/infra"
	"team-management/members/repository"
	"team-management/members/usecase"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger := infra.NewLoggerAdapter()

	config, err := infra.NewConfig()
	if err != nil {
		logger.Error("Failed to load .env file. Create one or set environment variables - ", err.Error())
	}

	context, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(context, options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	if err := mongoClient.Ping(context, nil); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	membersDBclient := infra.NewMongoDBAdapter(mongoClient, config.MongoDBDatabaseName, config.MongoDBCollectionName)
	membersRepository := repository.NewMembersRepository(membersDBclient)
	validator := infra.NewValidatorAdapter()

	createMemberUseCase := usecase.NewCreateMember(membersRepository)
	getMemberUseCase := usecase.NewGetMember(membersRepository)
	filterMemberUseCase := usecase.NewFilterMember(membersRepository)
	updateMemberUseCase := usecase.NewUpdateMember(membersRepository)
	deleteMemberUseCase := usecase.NewDeleteMember(membersRepository)

	handlers := &api.Handlers{
		ErrorHandler: api.NewErrorHandler(logger),
		CreateMember: api.NewCreateMember(createMemberUseCase, validator),
		GetMember:    api.NewGetMember(getMemberUseCase),
		FilterMember: api.NewFilterMember(filterMemberUseCase),
		UpdateMember: api.NewUpdateMember(updateMemberUseCase, validator),
		DeleteMember: api.NewDeleteMember(deleteMemberUseCase),
	}

	router := api.NewRouter(handlers)
	handler := router.Register(logger)

	logger.Info(fmt.Sprintf("starting %s webserver", config.AppName))
	if err := http.ListenAndServe(config.AppPort, handler); err != nil {
		logger.Error(err)
	}
}
