package controllers

import (
	"Store/src/employee/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllEmployeeController struct {
	GetAllEmployeesUseCase *application.GetAllEmployee
}

func NewGetAllEmployeeController(getAll *application.GetAllEmployee) *GetAllEmployeeController {
	return &GetAllEmployeeController{GetAllEmployeesUseCase: getAll}
}

func (ctrl *GetAllEmployeeController) GetAllEmployees(c *gin.Context) {
	employees, err := ctrl.GetAllEmployeesUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(employees) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay empleados disponibles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"employees": employees})
}
