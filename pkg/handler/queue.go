package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/types"
)

// @Summary Get Queue List
// @Tags queue
// @Description get all queue list
// @ID get-queue-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []types.QueueItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/queue [get]
func (h *Handler) getQueueLists(c *gin.Context) {
	items, err := h.services.Queue.GetQueueList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Add New Queue Item
// @Tags queue
// @Description add new ticket (item queue) in the end of the queue
// @ID add-new-ticket
// @Accept  json
// @Produce  json
// @Success 200 {object} types.QueueItemNumber
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/queue/service [get]
func (h *Handler) addQueueItem(c *gin.Context) {
	serviceType := c.Param("service")

	queueItemNumber, err := h.services.Queue.AddQueueItem(serviceType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, types.QueueItemNumber{Ticket: queueItemNumber})
}
