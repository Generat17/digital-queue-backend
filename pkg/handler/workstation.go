package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/types"
)

type getWorkstationListsResponse struct {
	Data []types.Workstation `json:"data"`
}

// @Summary Get Workstation Data
// @Tags workstation
// @Description get workstation data
// @ID get-workstation
// @Accept  json
// @Produce  json
// @Success 200 {object} getWorkstationListsResponse "response"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/workstation  [post]
func (h *Handler) getWorkstation(c *gin.Context) {
	items, err := h.services.Workstation.GetWorkstationList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorkstationListsResponse{Data: items})
}
