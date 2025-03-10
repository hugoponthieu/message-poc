package app

import (
	"message/controller"
	"message/repository"
	"message/router"
	"message/service"
	"message/services/mongo_client"
)

type App struct {
	messageRouter     router.MessageRouter
	messageService    service.MessageService
	messageController controller.MessageController
	messageRepository repository.MessageRepository
	mongo             mongo_client.MongoClient
}

func initApp() *App {
	mongo := mongo_client.NewMongoClient()

	messageRepository := mongo.NewMongoRepository()
	messageService := service.NewMessageService(messageRepository)
	messageController := controller.NewMessageController(messageService)
	messageRouter := router.NewMessageRouter(messageController)

	return &App{
		messageRouter:     messageRouter,
		messageService:    messageService,
		messageController: messageController,
		messageRepository: messageRepository,
		mongo:             mongo,
	}
}
