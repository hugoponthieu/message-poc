package app

import (
	"message/config"
	"message/controller"
	"message/infrastructure"
	"message/internal/api/http"
	"message/repository"
	"message/repository/mongo"
	"message/router"
	"message/service"
)


type App struct {
	config            config.AppConfig
	MessageRouter     router.MessageRouter
	MessageService    service.MessageService
	MessageController controller.MessageController
	MessageRepository repository.MessageRepository
	Mongo             infrastructure.MongoClient
	HttpServer        *http.Server
}

func InitApp(appConfig config.AppConfig) (*App, error) {
	mongo_client, err := infrastructure.NewMongoClient(appConfig.MongoConfig)
	if err != nil {
		return nil, err
	}

	messageRepository := mongo.NewMongoRepository(mongo_client.Db)
	messageService := service.NewMessageService(messageRepository)
	messageController := controller.NewMessageController(messageService)
	messageRouter := router.NewMessageRouter(messageController)

	httpServer := http.NewServer(appConfig.Port, appConfig.AllowOrigin)
	return &App{
		config:            appConfig,
		MessageRouter:     messageRouter,
		MessageService:    messageService,
		MessageController: messageController,
		MessageRepository: messageRepository,
		Mongo:             *mongo_client,
		HttpServer:        httpServer,
	}, nil
}

func (app *App) Start() error {
	
	app.MessageRouter.RegisterRoutes(&app.HttpServer.Engine.RouterGroup)
	app.HttpServer.Start()
	return nil
}
