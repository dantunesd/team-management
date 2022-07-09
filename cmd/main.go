package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	teammanagement "team-management"
	"team-management/members/infrastructure/logger"
	"team-management/members/infrastructure/mongodb"
	"team-management/members/infrastructure/repository"
	"team-management/members/infrastructure/validator"
	"team-management/members/presentation/api"
	"team-management/members/usecase"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger := logger.NewLogrusAdapter()

	config, err := teammanagement.NewConfig()
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

	membersDBclient := mongodb.NewMongoDBAdapter(mongoClient, config.MongoDBDatabaseName, config.MongoDBCollectionName)
	membersRepository := repository.NewMembersRepository(membersDBclient)
	validator := validator.NewValidatorAdapter()

	getMemberUseCase := usecase.NewGetMember(membersRepository)
	createMemberUseCase := usecase.NewCreateMember(membersRepository)
	filterMemberUseCase := usecase.NewFilterMember(membersRepository)
	updateMemberUseCase := usecase.NewUpdateMember(membersRepository)
	deleteMemberUseCase := usecase.NewDeleteMember(membersRepository)

	errorHandler := api.NewErrorHandler(logger)
	getMemberHandler := api.NewGetMember(getMemberUseCase)
	createMemberHandler := api.NewCreateMember(createMemberUseCase, validator)
	filterMemberHandler := api.NewFilterMember(filterMemberUseCase)
	updateMemberHandler := api.NewUpdateMember(updateMemberUseCase, validator)
	deleteMemberHandler := api.NewDeleteMember(deleteMemberUseCase)

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true}))

	router.Route("/members", func(c chi.Router) {
		c.Get("/", errorHandler.Handle(filterMemberHandler.Handle))
		c.Post("/", errorHandler.Handle(createMemberHandler.Handle))
		c.Get("/{id}", errorHandler.Handle(getMemberHandler.Handle))
		c.Put("/{id}", errorHandler.Handle(updateMemberHandler.Handle))
		c.Delete("/{id}", errorHandler.Handle(deleteMemberHandler.Handle))
	})

	logger.Info(fmt.Sprintf("starting %s webserver", config.AppName))
	if err := http.ListenAndServe(config.AppPort, router); err != nil {
		logger.Error(err)
	}
}
