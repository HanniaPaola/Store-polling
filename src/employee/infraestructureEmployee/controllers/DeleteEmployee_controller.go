package controllers

import (
	"Store/src/employee/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEmployeeController struct {
	DeleteEmployeeUseCase *application.DeleteEmployee
}

func NewDeleteEmployeeController(delete *application.DeleteEmployee) *DeleteEmployeeController {
	return &DeleteEmployeeController{DeleteEmployeeUseCase: delete}
}

func (ctrl *DeleteEmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.DeleteEmployeeUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empleado eliminado"})
}
