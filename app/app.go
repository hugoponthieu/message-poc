package app

import (
	"message/config"
	"message/controller"
	"message/repository"
	"message/repository/mongo"
	"message/router"
	"message/service"
	"message/services/mongo_client"

	"github.com/gin-gonic/gin"
)

type App struct {
	config            config.AppConfig
	MessageRouter     router.MessageRouter
	MessageService    service.MessageService
	MessageController controller.MessageController
	MessageRepository repository.MessageRepository
	Mongo             mongo_client.MongoClient
}

func InitApp(appConfig config.AppConfig) (*App, error) {
	mongo_client, err := mongo_client.NewMongoClient(appConfig.MongoConfig)
	if err != nil {
		return nil, err
	}

	messageRepository := mongo.NewMongoRepository(mongo_client.Db)
	messageService := service.NewMessageService(messageRepository)
	messageController := controller.NewMessageController(messageService)
	messageRouter := router.NewMessageRouter(messageController)

	return &App{
		config:            appConfig,
		MessageRouter:     messageRouter,
		MessageService:    messageService,
		MessageController: messageController,
		MessageRepository: messageRepository,
		Mongo:            *mongo_client,
	}, nil
}

func (app *App) Start() error {
	router := gin.Default()
	app.MessageRouter.RegisterRoutes(&router.RouterGroup)
	router.Run(":" + app.config.Port)
	return nil
}
