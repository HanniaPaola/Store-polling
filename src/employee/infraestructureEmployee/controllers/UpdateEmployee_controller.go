package controllers

import (
	"Store/src/employee/application"
	"Store/src/employee/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEmployeeController struct {
	UpdateEmployeeUseCase *application.UpdateEmployee
}

func NewUpdateEmployeeController(update *application.UpdateEmployee) *UpdateEmployeeController {
	return &UpdateEmployeeController{UpdateEmployeeUseCase: update}
}

func (ctrl *UpdateEmployeeController) UpdateEmployee(c *gin.Context) {
	var req struct {
		Name  string  `json:"name"`
		Position string `json:"position"`
		Salary int `json:"salary"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	employee := domain.NewEmployee(req.Name, req.Position, int32(req.Salary)) // Cambiar a NewEmployee
	employee.SetID(int32(id))

	err = ctrl.UpdateEmployeeUseCase.Run(*employee) // Cambiar a Run
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empleado actualizado"})
}
