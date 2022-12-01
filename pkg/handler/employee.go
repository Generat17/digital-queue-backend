package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/types"
	"strconv"
)

// @Summary Get Employee Lists
// @Tags employee
// @Description get list employee
// @ID get-employee-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} types.GetEmployeeListsResponse
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

	c.JSON(http.StatusOK, types.GetEmployeeListsResponse{Data: items})
}

// @Summary Get New Client
// @Security ApiKeyAuth
// @Tags client
// @Description get an available client from the queue
// @ID get-new-client
// @Accept  json
// @Produce  json
// @Success 200 {object} types.GetNewClientResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/employee/client [post]
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

	c.JSON(http.StatusOK, client)
}

type confirmClientInput struct {
	NumberQueue string `json:"numberQueue" binding:"required"`
}

// @Summary Confirm Client
// @Security ApiKeyAuth
// @Tags client
// @Description confirms that the client has approached the workstation
// @ID confirm-client
// @Accept  json
// @Produce  json
// @Param input body confirmClientInput true "credentials"
// @Success 200 {object} types.ConfirmClientResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/employee/confirmClient [post]
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

	c.JSON(http.StatusOK, types.ConfirmClientResponse{NumberQueue: client})
}

// @Summary End Client
// @Security ApiKeyAuth
// @Tags client
// @Description complete the client
// @ID end-client
// @Accept  json
// @Produce  json
// @Success 200 {object} types.ConfirmClientResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/employee/endClient [post]
func (h *Handler) endClient(c *gin.Context) {
	employeeId, _ := c.Get(userCtx)
	empId := employeeId.(int)

	client, err := h.services.Queue.EndClient(empId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, types.ConfirmClientResponse{NumberQueue: client})
}

// @Summary Get Employee Status
// @Security ApiKeyAuth
// @Tags employee
// @Description get the current status of an employee
// @ID get-status-employee
// @Accept  json
// @Produce  json
// @Success 200 {object} types.EmployeeStatusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/employee/getStatus [post]
func (h *Handler) getStatus(c *gin.Context) {
	employeeId, _ := c.Get(userCtx)
	empId := employeeId.(int)

	status, err := h.services.Authorization.GetStatusEmployee(empId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, types.EmployeeStatusResponse{EmployeeStatus: status})
}
