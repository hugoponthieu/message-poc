package router

import (
	"message/controller"

	"github.com/gin-gonic/gin"
)

type MessageRouter struct {
	controller controller.MessageController
}

func NewMessageRouter(controller controller.MessageController) MessageRouter {
	return MessageRouter{
		controller: controller,
	}
}

func (r *MessageRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/messages/:id", r.controller.GetMessage)
	// Mget
	router.POST("/messages/batch", r.controller.GetMessages)
	router.POST("/messages", r.controller.CreateMessage)
	router.PUT("/messages/:id", r.controller.UpdateMessage)
	router.DELETE("/messages/:id", r.controller.DeleteMessage)
	router.GET("/messages/search", r.controller.SearchMessages)
}
