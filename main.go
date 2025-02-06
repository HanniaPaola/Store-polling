package main

import (
	"log"

	"Store/src/core"
	"Store/src/polling"
	"Store/src/employee/infraestructureEmployee"
	"Store/src/employee/infraestructureEmployee/routesEmployee"
	"Store/src/products/infraestructure"
	"Store/src/products/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	mysqlConn, err := core.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error al conectar a MySQL: %v", err)
	}
	defer mysqlConn.Close()

	mysql := &infraestructure.MySQL{Conn: mysqlConn.Conn}                  
	mysql1 := &infraestructureEmployee.MySQLEmployee{Conn: mysqlConn.Conn} 
	r := gin.Default()

		// Rutas para short polling
		r.GET("/polling/products", func(c *gin.Context) {
			polling.ShortPolling(c, mysql)
		})
		r.POST("/polling/update", func(c *gin.Context) {
			polling.UpdateProduct(c, mysql)
		})
	
		// Ruta para long polling
		r.GET("/polling/long", polling.LongPolling)
	
		// Iniciar simulación de actualizaciones
		go polling.simulateUpdates()

	routes.RegisterProductRoutes(r, mysql)
	routesEmployee.RegisterEmployeeRoutes(r, mysql1)

	r.Run(":8080")
}
