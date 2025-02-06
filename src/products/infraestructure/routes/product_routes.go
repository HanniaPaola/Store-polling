package routes

import (
	"Store/src/products/application"
	"Store/src/products/infraestructure"
	"Store/src/products/infraestructure/controllers"
	
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, mysql *infraestructure.MySQL) {
	// Crear controlador de producto
	controllerCreate := controllers.NewCreateProductController(
		application.NewCreateProduct(mysql),
	)

	controllerGet := controllers.NewGetAllProductController(
		application.NewGetAllProduct(mysql),
	)

	controllerGetByID := controllers.NewGetProductByIDController(
		application.NewGetProductByID(mysql),
	)
	controllerUpdate := controllers.NewUpdateProductController(
		application.NewUpdateProduct(mysql),
	)
	controllerDelete := controllers.NewDeleteProductController(
		application.NewDeleteProduct(mysql),
	)

	r.POST("/products", controllerCreate.CreateProduct)
	r.GET("/products", controllerGet.GetAllProducts)
	r.GET("/products/:id", controllerGetByID.GetProductByID)
	r.PUT("/products/:id", controllerUpdate.UpdateProduct)
	r.DELETE("/products/:id", controllerDelete.DeleteProduct)
}
