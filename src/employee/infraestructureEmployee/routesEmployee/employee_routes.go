package routesEmployee

import (
	"Store/src/employee/application"
	"Store/src/employee/infraestructureEmployee"
	"Store/src/employee/infraestructureEmployee/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(r *gin.Engine, mysql *infraestructureEmployee.MySQLEmployee) {
	// Crear controlador de empleado
	controller := controllers.NewCreateEmployeeController(
		application.NewCreateEmployee(mysql),
	)

	controllerGet := controllers.NewGetAllEmployeeController(
		application.NewGetAllEmployee(mysql),
	)

	controllerGetByID := controllers.NewGetEmployeeByIDController(
		application.NewGetEmployeeByID(mysql),
	)
	controllerUpdate := controllers.NewUpdateEmployeeController(
		application.NewUpdateEmployee(mysql),
	)
	controllerDelete := controllers.NewDeleteEmployeeController(
		application.NewDeleteEmployee(mysql),
	)

	r.POST("/employee", controller.CreateEmployee)
	r.GET("/employee", controllerGet.GetAllEmployees)
	r.GET("/employee/:id", controllerGetByID.GetEmployeeByID)
	r.PUT("/employee/:id", controllerUpdate.UpdateEmployee)
	r.DELETE("/employee/:id", controllerDelete.DeleteEmployee)
}
