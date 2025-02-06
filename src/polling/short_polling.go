package polling

import (
	"net/http"
	"Store/src/products/infraestructure"

	"github.com/gin-gonic/gin"
)

// Producto representará la estructura de un producto
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Endpoint de short polling para obtener productos
func ShortPolling(c *gin.Context, mysql *infraestructure.MySQL) {
	// Obtén todos los productos
	products, err := mysql.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Endpoint para actualizar un producto
func UpdateProduct(c *gin.Context, mysql *infraestructure.MySQL) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mysql.Update(product.ID, product.Name, product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
