package handler

import (
	"net/http"
	"server/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
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
func (h *Handler) signUp(c *gin.Context) {
	var input types.Employee

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateEmployee(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description username
// @ID username
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signInWorkstation(c *gin.Context) {
	var input signInInput
	workstationId, _ := strconv.Atoi(c.Param("workstation"))

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.services.Authorization.GenerateTokenWorkstation(input.Username, input.Password, workstationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.Authorization.GenerateRefreshToken()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	employee, err := h.services.Authorization.GetEmployee(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.services.Authorization.SetSession(refreshToken, workstationId, employee.EmployeeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"employee":     employee,
		"workstation":  workstationId,
	})
}

type refreshTokenInput struct {
	WorkstationId string `json:"workstationId" binding:"required"`
	EmployeeId    string `json:"employeeId" binding:"required"`
	RefreshToken  string `json:"refreshToken" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description username
// @ID username
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) refresh(c *gin.Context) {

	var input refreshTokenInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	WorkstationId, _ := strconv.Atoi(input.WorkstationId)
	EmployeeId, _ := strconv.Atoi(input.EmployeeId)

	accessToken, err := h.services.Authorization.UpdateTokenWorkstation(EmployeeId, WorkstationId, input.RefreshToken)
	if err != nil {
		if accessToken == "refreshToken is invalid" {
			newErrorResponse(c, http.StatusForbidden, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	refreshToken, err := h.services.Authorization.GenerateRefreshToken()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	employee, err := h.services.Authorization.GetEmployeeById(EmployeeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.services.Authorization.SetSession(refreshToken, WorkstationId, EmployeeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"employee":     employee,
		"workstation":  WorkstationId,
	})
}

type logoutInput struct {
	EmployeeId string `json:"employeeId" binding:"required"`
}

// @Summary SignUp
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
func (h *Handler) logout(c *gin.Context) {
	var input logoutInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	employeeId, _ := strconv.Atoi(input.EmployeeId)

	res, err := h.services.Authorization.LogOut(employeeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
