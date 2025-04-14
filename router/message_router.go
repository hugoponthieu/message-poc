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
	messageRouter := router.Group("/messages")
	messageRouter.GET("/:id", r.controller.GetMessage)
	// Mget
	messageRouter.POST("/batch", r.controller.GetMessages)
	messageRouter.POST("", r.controller.CreateMessage)
	messageRouter.PUT("/:id", r.controller.UpdateMessage)
	messageRouter.DELETE("/:id", r.controller.DeleteMessage)
	messageRouter.GET("/search", r.controller.SearchMessages)
}
