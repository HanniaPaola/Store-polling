package controllers

import (
	"Store/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	DeleteProductUseCase *application.DeleteProduct
}

func NewDeleteProductController(delete *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{DeleteProductUseCase: delete}
}

func (ctrl *DeleteProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.DeleteProductUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}