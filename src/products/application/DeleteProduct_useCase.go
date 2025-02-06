package application

import "Store/src/products/domain"

type DeleteProduct struct {
	repository domain.IProduct
}

func NewDeleteProduct(repository domain.IProduct) *DeleteProduct {
	return &DeleteProduct{repository}
}

func (d *DeleteProduct) Run(id int32) error {
	return d.repository.Delete(id)
}
