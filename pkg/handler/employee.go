package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/types"
	"strconv"
)

type getEmployeeListsResponse struct {
	Data []types.Employee `json:"data"`
}

// @Summary Get Employee Lists
// @Security ApiKeyAuth
// @Tags employee
// @Description get all data about all employee
// @ID get-employee-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getEmployeeListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/employee [get]
func (h *Handler) getEmployeeLists(c *gin.Context) {
	items, err := h.services.Employee.GetEmployeeList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getEmployeeListsResponse{Data: items})
}

func (h *Handler) getNewClient(c *gin.Context) {
	employeeId, _ := c.Get(userCtx)
	empId := employeeId.(int)
	workstationId, _ := c.Get(workstationCtx)
	workId := workstationId.(int)

	client, err := h.services.Queue.GetNewClient(empId, workId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Print(client)

	c.JSON(http.StatusOK, client)
}

type confirmClientInput struct {
	NumberQueue string `json:"numberQueue" binding:"required"`
}

func (h *Handler) confirmClient(c *gin.Context) {
	employeeId, _ := c.Get(userCtx)
	empId := employeeId.(int)

	var input confirmClientInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	numberQueue, _ := strconv.Atoi(input.NumberQueue)

	client, err := h.services.Queue.ConfirmClient(numberQueue, empId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) endClient(c *gin.Context) {
	employeeId, _ := c.Get(userCtx)
	empId := employeeId.(int)

	client, err := h.services.Queue.EndClient(empId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) getStatus(c *gin.Context) {
	employeeId, _ := c.Get(userCtx)
	empId := employeeId.(int)

	status, err := h.services.Authorization.GetStatusEmployee(empId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status)
}
