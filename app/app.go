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
	messageRouter     router.MessageRouter
	messageService    service.MessageService
	messageController controller.MessageController
	messageRepository repository.MessageRepository
	mongo             mongo_client.MongoClient
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
		messageRouter:     messageRouter,
		messageService:    messageService,
		messageController: messageController,
		messageRepository: messageRepository,
		mongo:             *mongo_client,
	}, nil
}

func (app *App) Start() error {
	router := gin.Default()
	app.messageRouter.RegisterRoutes(&router.RouterGroup)
	router.Run(":" + app.config.Port)
	return nil
}
