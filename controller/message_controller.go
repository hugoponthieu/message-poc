package controller

import (
	"message/service"
	"message/types/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	service service.MessageService
}

func NewMessageController(messageService service.MessageService) MessageController {
	return MessageController{
		service: messageService,
	}
}

// GetMessage godoc
// @Summary Get a message by ID
// @Description Get a single message by its ID
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} message.Message
// @Failure 404 {object} ErrorResponse
// @Router /messages/{id} [get]
func (c *MessageController) GetMessage(ctx *gin.Context) {
	println("message")
	id := ctx.Param("id")
	msg, err := c.service.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, msg)
}

// GetMessages godoc
// @Summary Get multiple messages by IDs
// @Description Get multiple messages by their IDs
// @Tags messages
// @Accept json
// @Produce json
// @Param ids body []string true "Message IDs"
// @Success 200 {object} types.MGetResult
// @Router /messages/batch [post]
func (c *MessageController) GetMessages(ctx *gin.Context) {
	var ids []string
	if err := ctx.BindJSON(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.Mget(ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// CreateMessage godoc
// @Summary Create a new message
// @Description Create a new message
// @Tags messages
// @Accept json
// @Produce json
// @Param message body message.Message true "Message object"
// @Success 201 {object} message.Message
// @Router /messages [post]
func (c *MessageController) CreateMessage(ctx *gin.Context) {
	var msg message.Message
	if err := ctx.BindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.service.Create(&msg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

// UpdateMessage godoc
// @Summary Update a message
// @Description Update an existing message
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Param message body message.UpdateMessage true "Updated message object"
// @Success 200 {object} message.Message
// @Router /messages/{id} [put]
func (c *MessageController) UpdateMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateMsg message.UpdateMessage
	if err := ctx.BindJSON(&updateMsg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.service.Update(id, &updateMsg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

// DeleteMessage godoc
// @Summary Delete a message
// @Description Delete a message by its ID
// @Tags messages
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 204 "No Content"
// @Router /messages/{id} [delete]
func (c *MessageController) DeleteMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GetMessagesByChannel godoc
// @Summary Get messages by channel
// @Description Get messages for a specific channel with pagination
// @Tags messages
// @Accept json
// @Produce json
// @Param channelId path string true "Channel ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} message.Message
// @Router /messages/channel/{channelId} [get]
func (c *MessageController) GetMessagesByChannel(ctx *gin.Context) {
	channelID := ctx.Param("channelId")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	messages, err := c.service.GetByChannel(channelID, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

// SearchMessages godoc
// @Summary Search messages
// @Description Search messages with various filters
// @Tags messages
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Param channelId query string false "Channel ID"
// @Param serverId query string false "Server ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} message.Message
// @Router /messages/search [get]
func (c *MessageController) SearchMessages(ctx *gin.Context) {
	query := ctx.Query("query")
	channelID := ctx.Query("channelId")
	serverID := ctx.Query("serverId")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	var channelIDPtr, serverIDPtr *string
	if channelID != "" {
		channelIDPtr = &channelID
	}
	if serverID != "" {
		serverIDPtr = &serverID
	}

	messages, err := c.service.Search(query, channelIDPtr, serverIDPtr, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}
