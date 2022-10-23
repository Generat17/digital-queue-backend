package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get All Responsibility
// @Tags responsibility
// @Description get all data about responsibility
// @ID get-responsibility-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} []types.Responsibility
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/responsibility [get]
func (h *Handler) getResponsibilityList(c *gin.Context) {
	items, err := h.services.Responsibility.GetResponsibilityList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
