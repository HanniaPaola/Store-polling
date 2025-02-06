package controllers

import (
	"Store/src/employee/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEmployeeByIDController struct {
	GetEmployeeByIDUseCase *application.GetEmployeeByID
}

func NewGetEmployeeByIDController(getByID *application.GetEmployeeByID) *GetEmployeeByIDController {
	return &GetEmployeeByIDController{GetEmployeeByIDUseCase: getByID}
}

func (ctrl *GetEmployeeByIDController) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	employee, err := ctrl.GetEmployeeByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"employee": employee.ToJSON()})
}
