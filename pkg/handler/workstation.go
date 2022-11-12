package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/types"
)

type getWorkstationListsResponse struct {
	Data []types.Workstation `json:"data"`
}

// @Summary Workstation
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) getWorkstation(c *gin.Context) {
	items, err := h.services.Workstation.GetWorkstationList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorkstationListsResponse{Data: items})
}
