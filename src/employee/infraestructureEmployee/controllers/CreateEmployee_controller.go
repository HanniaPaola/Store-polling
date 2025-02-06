package controllers

import (
	"Store/src/employee/application"
	"Store/src/employee/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeController struct {
	CreateEmployeeUseCase *application.CreateEmployee
}

func NewCreateEmployeeController(create *application.CreateEmployee) *CreateEmployeeController {
	return &CreateEmployeeController{CreateEmployeeUseCase: create}
}

func (ctrl *CreateEmployeeController) CreateEmployee(c *gin.Context) {
	var req struct {
		Name  string  `json:"name"`
		Position string `json:"position"`
		Salary int `json:"salary"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	employee := domain.NewEmployee(req.Name, req.Position, int32(req.Salary))
	err := ctrl.CreateEmployeeUseCase.Run(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Empleado creado", "employee": employee.ToJSON()})
}
